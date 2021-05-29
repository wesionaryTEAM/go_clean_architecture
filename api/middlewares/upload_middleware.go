package middlewares

import (
	"clean-architecture/api/responses"
	"clean-architecture/infrastructure"
	"clean-architecture/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Extension string

const (
	JPEGFile Extension = "jpeg"
)

type UploadConfig struct {
	// FieldName where to pull multipart files from
	FieldName string

	// Extensions array of extensions
	Extensions []Extension
}

type UploadMiddleware struct {
	logger infrastructure.Logger
	bucket services.BucketService
	config UploadConfig
}

func NewUploadMiddleware(
	logger infrastructure.Logger,
	bucket services.BucketService,
) UploadMiddleware {
	return UploadMiddleware{
		config: UploadConfig{
			FieldName:  "file",
			Extensions: []Extension{JPEGFile},
		},
		bucket: bucket,
		logger: logger,
	}
}

// Config sets up the file upload config
func (u UploadMiddleware) Config(config UploadConfig) UploadMiddleware {
	u.config = config
	return u
}

// Handle handles file upload
func (u UploadMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, fileHeader, _ := c.Request.FormFile(u.config.FieldName)
		if file != nil && fileHeader != nil {
			url, err := u.bucket.UploadFile(c.Request.Context(), file, fileHeader.Filename, fileHeader.Filename)
			if err != nil {
				u.logger.Error("file-upload-error: ", err.Error())
				responses.ErrorJSON(c, http.StatusInternalServerError, err)
				c.Abort()
				return
			}
			c.Set(u.config.FieldName, url)
		}
	}
}

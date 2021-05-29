package middlewares

import (
	"bytes"
	"clean-architecture/api/responses"
	"clean-architecture/infrastructure"
	"clean-architecture/services"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"golang.org/x/sync/errgroup"
)

type Extension string

const (
	JPEGFile Extension = ".jpeg"
	JPGFile  Extension = ".jpg"
	PNGFile  Extension = ".png"
)

type UploadConfig struct {
	// FieldName where to pull multipart files from
	FieldName string

	// BucketFolder where to put the uploaded files to
	BucketFolder string

	// Extensions array of extensions
	Extensions []Extension

	// ThumbnailEnabled set whether thumbnail is enabled or nor
	ThumbnailEnabled bool

	// ThumbnailWidth set thumbnail width
	ThumbnailWidth uint
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
			FieldName:        "file",
			BucketFolder:     "",
			Extensions:       []Extension{JPEGFile, PNGFile, JPGFile},
			ThumbnailEnabled: true,
			ThumbnailWidth:   100,
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

			ext := filepath.Ext(fileHeader.Filename)
			if !u.matchesExtension(ext) {
				u.logger.Error("file-upload-error: ", "extension mismatch")
				responses.ErrorJSON(c, http.StatusBadRequest, "file extension not supported")
				c.Abort()
				return
			}

			uploadFileName, fileUID := u.randomFileName(ext)
			errorGroup, ctx := errgroup.WithContext(c.Request.Context())

			errorGroup.Go(func() error {
				url, err := u.bucket.UploadFile(ctx, file, uploadFileName, fileHeader.Filename)
				c.Set(u.config.FieldName, url)
				return err
			})

			if u.config.ThumbnailEnabled {
				errorGroup.Go(func() error {

					e := Extension(ext)
					properExtension := e == JPEGFile || e == JPGFile || e == PNGFile

					if !properExtension {
						return errors.New("not proper extension for thumbnail generation")
					}

					img, err := u.createThumbnail(file, ext)
					if err != nil {
						return err
					}

					resizeFileName := u.bucketPath(fmt.Sprintf("%s_thumb%s", fileUID, ext))
					_, err = u.bucket.UploadFile(ctx, img, resizeFileName, fileHeader.Filename)
					if err != nil {
						return err
					}
					return nil
				})
			}

			if err := errorGroup.Wait(); err != nil {
				u.logger.Error("file-upload-error: ", err.Error())
				responses.ErrorJSON(c, http.StatusInternalServerError, err)
				c.Abort()
				return
			}

		}
		c.Next()
	}
}

func (u UploadMiddleware) matchesExtension(ext string) bool {
	for _, e := range u.config.Extensions {
		if e == Extension(ext) {
			return true
		}
	}
	return false
}

func (u UploadMiddleware) randomFileName(ext string) (string, string) {
	randUUID, _ := uuid.NewRandom()
	fileName := randUUID.String() + ext
	return u.bucketPath(fileName), randUUID.String()
}

func (u UploadMiddleware) bucketPath(fileName string) string {
	if u.config.BucketFolder != "" {
		return fmt.Sprintf("%s/%s", u.config.BucketFolder, fileName)
	}
	return fileName
}

// createThumbnail creates thumbnail from multipart file
func (u UploadMiddleware) createThumbnail(file multipart.File, ext string) (*bytes.Buffer, error) {
	var img image.Image
	var err error

	file.Seek(0, 0)
	if Extension(ext) == JPGFile || Extension(ext) == JPEGFile {
		img, err = jpeg.Decode(file)
	}
	if Extension(ext) == PNGFile {
		img, err = png.Decode(file)
	}

	if err != nil {
		return nil, err
	}

	resizeImage := resize.Resize(u.config.ThumbnailWidth, 0, img, resize.Lanczos3)
	buff := new(bytes.Buffer)
	if Extension(ext) == JPGFile || Extension(ext) == JPEGFile {
		if err := jpeg.Encode(buff, resizeImage, nil); err != nil {
			return nil, err
		}
	}
	if Extension(ext) == PNGFile {
		if err := png.Encode(buff, resizeImage); err != nil {
			return nil, err
		}
	}

	return buff, nil
}

package responses

import (
	"clean-architecture/pkg/errorz"
	"clean-architecture/pkg/framework"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, status int, data any, meta map[string]any) {
	body := gin.H{"success": true, "data": data}
	if len(meta) > 0 {
		body["meta"] = meta
	}
	c.JSON(status, body)
}

func PaginationSuccess(c *gin.Context, status int, data any, total int64) {
	limit, _ := c.MustGet(framework.Limit).(int)
	page, _ := c.MustGet(framework.Page).(int)
	hasNext := (total - int64(limit)*int64(page)) > 0
	meta := map[string]any{"pagination": map[string]any{"page": page, "limit": limit, "total": total, "has_next": hasNext}}
	Success(c, status, data, meta)
}

// Error standardized error response
func Error(c *gin.Context, api *errorz.APIError) {
	if api == nil {
		api = errorz.ErrInternal
	}
	body := gin.H{"success": false, "code": api.Code, "message": api.Message}
	if len(api.Details) > 0 {
		body["details"] = api.Details
	}
	c.JSON(api.StatusCode, body)
}

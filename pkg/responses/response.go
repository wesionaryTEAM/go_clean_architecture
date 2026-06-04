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

// ErrorBody is the serialized form of an APIError.
// APIError.StatusCode (transport) and APIError.Cause (internal) are intentionally omitted.
type ErrorBody struct {
	Code     string          `json:"code"`
	Severity errorz.Severity `json:"severity"`
	Params   map[string]any  `json:"params,omitempty"`
	Message  string          `json:"message,omitempty"`
}

// ErrorResponse is the standardized error envelope.
type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

// Error writes a standardized error response.
func Error(c *gin.Context, api *errorz.APIError) {
	if api == nil {
		api = errorz.ErrInternal
	}
	c.JSON(api.StatusCode, ErrorResponse{Error: ErrorBody{
		Code:     api.Code,
		Severity: api.Severity,
		Params:   api.Params,
		Message:  api.Message,
	}})
}

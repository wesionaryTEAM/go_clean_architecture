package responses

import (
	"clean-architecture/pkg/framework"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field     string `json:"field"`
	FullField string `json:"full_field"`
	Message   string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "max":
		return "Should be less than " + fe.Param()
	case "min":
		return "Should be greater than " + fe.Param()
	case "email|e164":
		return "Invalid contact address format"
	}
	return "Unknown error"
}

func HandleValidationError(logger framework.Logger, c *gin.Context, err error) {
	logger.Error("VALIDATION ERROR:", err)
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{
				fe.Field(),
				fe.Namespace(),
				getErrorMsg(fe),
			}
		}
		ErrorJSON(c, http.StatusBadRequest, out)
	} else {
		ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
}

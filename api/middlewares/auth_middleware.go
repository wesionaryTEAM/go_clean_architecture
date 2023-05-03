package middlewares

import (
	"github.com/gin-gonic/gin"
)

// AuthMiddleware interface has basic methods to authenticate and autorize user
// Can be updated as per project requirements
type AuthMiddleware interface {
	HandleAuthWithRole(roles ...string) gin.HandlerFunc
}

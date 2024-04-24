package interfaces

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	HandleAuthWithRole(roles ...string) gin.HandlerFunc
}

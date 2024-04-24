package interfaces

import "github.com/gin-gonic/gin"

type PaginationMiddleware interface {
	Handle() gin.HandlerFunc
}

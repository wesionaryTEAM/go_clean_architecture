package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	GET(uri string, f func(c *gin.Context))
	POST(uri string, f func(c *gin.Context))
	SERVE(port string)
}

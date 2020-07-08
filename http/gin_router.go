package router

import (
	"github.com/gin-gonic/gin"
)

type ginRouter struct{}

var (
	ginDispatcher = gin.Default()
)

func NewGinRouter() Router {
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, f func(c *gin.Context)) {
	ginDispatcher.GET(uri, f)
}

func (*ginRouter) POST(uri string, f func(c *gin.Context)) {
	ginDispatcher.POST(uri, f)
}

func (*ginRouter) SERVE(port string) {
	_ = ginDispatcher.Run(port)
}

func (*ginRouter) GROUP(uri string) *gin.RouterGroup {
	return ginDispatcher.Group(uri)
}

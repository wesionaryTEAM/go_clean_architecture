package main

import (
	"net/http"

	"prototype2/repository"

	"github.com/gin-gonic/gin"

	router "prototype2/http"
)

var (
	httpRouter router.Router = router.NewGinRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	httpRouter.GET("/users", repository.GetUsers)
	httpRouter.POST("/users", repository.AddUser)

	httpRouter.SERVE(port)

}

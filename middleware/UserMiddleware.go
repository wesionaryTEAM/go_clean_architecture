package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserAuthenticationMiddleware(c *gin.Context) {
	fmt.Println("This is into the middle ware")
}
// package router

// import (
// 

// 	"github.com/gin-gonic/gin"
// )

// type Router interface {
// 	GET(uri string, f func(c *gin.Context))
// 	POST(uri string, f func(c *gin.Context))
// 	SERVE(port string)
// }


package router

import "net/http"

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

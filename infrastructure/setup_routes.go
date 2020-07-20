package infrastructure

import (
	"net/http"
	"os"

	"prototype2/api/routes"
	router "prototype2/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB, fb *auth.Client) {
	httpRouter := router.NewGinRouter()

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	// POST Routes
	routes.PostRoutes(httpRouter.GROUP("/posts"), db)

	// User Routes
	routes.UserRoutes(httpRouter.GROUP("/users"), db, fb)

	// Run server
	port := os.Getenv("SERVER_PORT")
	httpRouter.SERVE(port)
	if port == "" {
		httpRouter.SERVE(":8000")
	} else {
		httpRouter.SERVE(port)
	}
}

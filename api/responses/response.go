package responses

import (
	"github.com/gin-gonic/gin"
)

// JSON : json response function
func JSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"data": data})
}

// ErrorJSON : json error response function
func ErrorJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"error": data})
}

// SuccessJSON : json error response function
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"msg": data})
}

// JSONCount : json response function
func JSONCount(c *gin.Context, statusCode int, data interface{}, count int) {
	c.JSON(statusCode, gin.H{"data": data, "count": count})
}

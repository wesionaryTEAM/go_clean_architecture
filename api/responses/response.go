package responses

import (
	"clean-architecture/constants"

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

// JSONWithPagination : json response function
func JSONWithPagination(c *gin.Context, statusCode int, response map[string]interface{}) {
	limit, _ := c.MustGet(constants.Limit).(int64)
	size, _ := c.MustGet(constants.Page).(int64)

	c.JSON(
		200,
		gin.H{
			"data":       response["data"],
			"pagination": gin.H{"has_next": (response["count"].(int64) - limit*size) > 0, "count": response["count"]},
		})
}

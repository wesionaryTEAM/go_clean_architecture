package utils

import (
	"clean-architecture/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limit, _ := c.MustGet(constants.Limit).(int64)
		page, _ := c.MustGet(constants.Page).(int64)

		offset := (page - 1) * limit

		return db.Offset(int(offset)).Limit(int(limit))
	}
}

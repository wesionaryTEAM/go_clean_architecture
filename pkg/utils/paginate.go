package utils

import (
	"clean-architecture/pkg/framework"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limit, _ := c.MustGet(framework.Limit).(int64)
		page, _ := c.MustGet(framework.Page).(int64)

		offset := (page - 1) * limit

		return db.Offset(int(offset)).Limit(int(limit))
	}
}

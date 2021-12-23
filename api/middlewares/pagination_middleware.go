package middlewares

import (
	"clean-architecture/constants"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationMiddleware struct {
}

func NewPaginationMiddleware() PaginationMiddleware {
	return PaginationMiddleware{}
}

func (p PaginationMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		per_page, err := strconv.ParseInt(c.Query("per_page"), 10, 0)
		if err != nil {
			per_page = 10
		}

		page, err := strconv.ParseInt(c.Query("page"), 10, 0)
		if err != nil {
			page = 0
		}

		c.Set(constants.Limit, per_page)
		c.Set(constants.Page, page)

		c.Next()
	}
}

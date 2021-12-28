package middlewares

import (
	"clean-architecture/constants"
	"clean-architecture/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationMiddleware struct {
	logger lib.Logger
}

func NewPaginationMiddleware(logger lib.Logger) PaginationMiddleware {
	return PaginationMiddleware{logger: logger}
}

func (p PaginationMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.logger.Info("setting up pagination middleware")

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

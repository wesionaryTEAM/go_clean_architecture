package middlewares

import (
	"clean-architecture/pkg/framework"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationMiddleware struct {
	logger framework.Logger
}

func NewPaginationMiddleware(logger framework.Logger) PaginationMiddleware {
	return PaginationMiddleware{logger: logger}
}

func (p PaginationMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.logger.Info("setting up pagination middleware")

		perPage, err := strconv.ParseInt(c.Query("per_page"), 10, 0)
		if err != nil {
			perPage = 10
		}

		page, err := strconv.ParseInt(c.Query("page"), 10, 0)
		if err != nil {
			page = 0
		}

		c.Set(framework.Limit, perPage)
		c.Set(framework.Page, page)

		c.Next()
	}
}

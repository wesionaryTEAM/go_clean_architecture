package mocks

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/types"

	"github.com/gin-gonic/gin"
)

func MockPaginationHandler(c *gin.Context) {
	c.Set(framework.Limit, int64(10))
	c.Set(framework.Page, int64(1))
	c.Next()
}

func MockAuthSuccessHandler(c *gin.Context) {
	c.Set(framework.UID, types.BinaryUUID{}.String())
	c.Set(framework.Claims, map[string]interface{}{})
	c.Next()
}

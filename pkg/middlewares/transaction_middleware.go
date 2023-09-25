package middlewares

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DBTransactionMiddleware -> struct for transaction
type DBTransactionMiddleware struct {
	logger framework.Logger
	db     infrastructure.Database
}

// NewDBTransactionMiddleware -> new instance of transaction
func NewDBTransactionMiddleware(
	logger framework.Logger,
	db infrastructure.Database,
) DBTransactionMiddleware {
	return DBTransactionMiddleware{
		logger: logger,
		db:     db,
	}
}

// Handle -> It setup the database transaction middleware
func (m DBTransactionMiddleware) Handle() gin.HandlerFunc {
	m.logger.Info("setting up database transaction middleware")

	return func(c *gin.Context) {
		txHandle := m.db.DB.Begin()
		m.logger.Info("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set(framework.DBTransaction, txHandle)
		c.Next()

		if utils.StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			m.logger.Info("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				m.logger.Error("trx commit error: ", err)
			}
		} else {
			m.logger.Info("rolling back transaction due to status code: ", c.Writer.Status())
			txHandle.Rollback()
		}
	}
}

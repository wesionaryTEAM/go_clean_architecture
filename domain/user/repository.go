package user

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"

	"gorm.io/gorm"
)

// UserRepository database structure
type Repository struct {
	infrastructure.Database
	logger framework.Logger
}

// NewUserRepository creates a new user repository
func NewRepository(db infrastructure.Database, logger framework.Logger) Repository {
	return Repository{db, logger}
}

// WithTrx delegate transaction from user repository
func (r Repository) WithTrx(trxHandle *gorm.DB) Repository {
	if trxHandle != nil {
		r.logger.Debug("using WithTrx as trxHandle is not nil")
		r.Database.DB = trxHandle
	}
	return r
}

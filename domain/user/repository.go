package user

import (
	"clean-architecture/domain/models"
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

// For AutoMigrating (used in fx.Invoke)
func Migrate(r Repository) error {
	r.logger.Info("[Migrating...User]")
	if err := r.DB.AutoMigrate(&models.User{}); err != nil {
		r.logger.Error("[Migration failed...User]")
		return err
	}
	return nil
}

func (r *Repository) Create(user *models.User) error {
	r.logger.Info("[UserRepository...Create]")
	return r.DB.Create(&user).Error
}

// WithTrx delegate transaction from user repository
func (r Repository) WithTrx(trxHandle *gorm.DB) Repository {
	if trxHandle != nil {
		r.logger.Debug("using WithTrx as trxHandle is not nil")
		r.Database.DB = trxHandle
	}
	return r
}

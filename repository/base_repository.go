package repository

import (
	"clean-architecture/infrastructure"

	"gorm.io/gorm"
)

type BaseRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewBaseRepository(
	db infrastructure.Database,
	logger infrastructure.Logger,
) BaseRepository {
	return BaseRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx enables repository with transaction
func (b BaseRepository) WithTrx(trxHandle *gorm.DB) BaseRepository {
	if trxHandle == nil {
		b.logger.Zap.Error("Transaction Database not found in gin context. ")
		return b
	}
	b.db.DB = trxHandle
	return b
}

// Create the data
func (b BaseRepository) Create(m interface{}) error {
	b.logger.Zap.Info("create of base repository")
	return b.db.Create(m).Error
}

// Update updates user
func (b BaseRepository) Update(m interface{}) error {
	return b.db.Save(m).Error
}

// GetOne gets ont user
func (b BaseRepository) GetOne(m interface{}, pk interface{}) (err error) {
	return b.db.First(m, pk).Error
}

// Delete deletes the row of data
func (b BaseRepository) Delete(m interface{}, pk interface{}) error {
	return b.db.Delete(m, pk).Error
}

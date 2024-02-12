package mocks

import (
	"clean-architecture/pkg/infrastructure"

	"github.com/DATA-DOG/go-sqlmock"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBOutput struct {
	fx.Out
	DB   *infrastructure.Database
	Mock *sqlmock.Sqlmock
}

func NewMockDB() (DBOutput, error) {
	_db, mock, err := sqlmock.New()
	if err != nil {
		return DBOutput{}, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      _db,
		SkipInitializeWithVersion: true,
	}))

	if err != nil {
		return DBOutput{}, err
	}

	return DBOutput{
		DB: &infrastructure.Database{
			DB: db,
		},
		Mock: &mock,
	}, nil
}

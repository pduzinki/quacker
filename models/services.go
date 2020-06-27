package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Services struct {
	db *gorm.DB
	Us UserService
}

func NewServices(dialect, connectionInfo string) *Services {
	db, err := gorm.Open(dialect, connectionInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	return &Services{
		db: db,
		Us: NewUserService(db),
	}
}

func (s *Services) Close() error {
	return s.db.Close()
}

func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}).Error
}

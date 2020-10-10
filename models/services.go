package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres driver
)

// Services wraps all model services
type Services struct {
	db *gorm.DB
	Us UserService
	Qs QuackService
	Fs FollowService
	Hs HashtagService
}

// NewServices creates Services instance
func NewServices(dialect, connectionInfo, passwordPepper, hmacKey string) *Services {
	db, err := gorm.Open(dialect, connectionInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	return &Services{
		db: db,
		Us: NewUserService(db, passwordPepper, hmacKey),
		Qs: NewQuackService(db),
		Fs: NewFollowService(db),
		Hs: NewHashtagService(db),
	}
}

// Close closes database connection
func (s *Services) Close() error {
	return s.db.Close()
}

// AutoMigrate performs auto migration for dabatase models
func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}, &Quack{}, &Follow{}, &Hashtag{}).Error
}

// RebuildDatabase drops all current database tables and performs auto migration
func (s *Services) RebuildDatabase() error {
	err := s.db.DropTableIfExists(&User{}, &Quack{}, &Follow{}, &Hashtag{}).Error
	if err != nil {
		return err
	}
	return s.AutoMigrate()
}

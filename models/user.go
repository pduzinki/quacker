package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User ...
type User struct {
	gorm.Model
	Username          string `gorm:"not null;unique_index"`
	Email             string `gorm:"not null;unique_index"`
	Password          string `gorm:"-"`
	PasswordHash      string `gorm:"not null"`
	RememberToken     string `gorm:"-"`
	RememberTokenHash string `gorm:"not null;unique_index"`
}

// UserDB ...
type UserDB interface {
}

// UserService ...
type UserService interface {
	UserDB
}

type userService struct {
	db *gorm.DB
}

func NewUserService(connectionInfo string) UserService {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		panic(err)
	}

	return &userService{
		db: db,
	}
}

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres driver
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

// NewUserService creates UserService instance
func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

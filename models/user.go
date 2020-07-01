package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres driver
)

// User represents user data in the database
type User struct {
	gorm.Model
	Username          string `gorm:"not null;unique_index"`
	Email             string `gorm:"not null;unique_index"`
	Password          string `gorm:"-"`
	PasswordHash      string `gorm:"not null"`
	RememberToken     string `gorm:"-"`
	RememberTokenHash string `gorm:"not null;unique_index"`
}

// UserDB is an interface for interacting with user data in the database
type UserDB interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

// UserService is an interface for interacting with user model
type UserService interface {
	UserDB
}

type userService struct {
	UserDB
}

// NewUserService creates UserService instance
func NewUserService(db *gorm.DB) UserService {
	ug := userGorm{
		db: db,
	}

	uv := userValidator{
		&ug,
	}

	return &userService{
		&uv,
	}
}

type userValidator struct {
	UserDB
}

func (uv *userValidator) Create(user *User) error {
	// TODO

	return uv.UserDB.Create(user)
}

func (uv *userValidator) Update(user *User) error {
	// TODO

	return uv.UserDB.Update(user)
}

func (uv *userValidator) Delete(id uint) error {
	// TODO

	return uv.UserDB.Delete(id)
}

type userGorm struct {
	db *gorm.DB
}

func (ug *userGorm) Create(user *User) error {
	return ug.db.Create(user).Error
}

func (ug *userGorm) Update(user *User) error {
	return ug.db.Save(user).Error
}

func (ug *userGorm) Delete(id uint) error {
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error
}

package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres driver
)

// User represents user data in the database
type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique_index"`
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	// RememberToken     string `gorm:"-"`
	// RememberTokenHash string `gorm:"not null;unique_index"`
}

// UserDB is an interface for interacting with user data in the database
type UserDB interface {
	FindByUsername(username string) (*User, error)
	FindByEmail(email string) (*User, error)

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
	ug := newUserGorm(db)
	uv := newUserValidator(ug)

	return &userService{
		uv,
	}
}

type userValidator struct {
	UserDB
	EmailRegex *regexp.Regexp
}

func newUserValidator(u UserDB) *userValidator {
	return &userValidator{
		UserDB:     u,
		EmailRegex: regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
	}
}

func (uv *userValidator) FindByUsername(username string) (*User, error) {
	u := User{}
	u.Username = username
	err := runUserValidatorFuncs(&u, uv.usernameNormalize)
	if err != nil {
		return nil, err
	}

	return uv.UserDB.FindByUsername(username)
}

func (uv *userValidator) FindByEmail(email string) (*User, error) {
	u := User{}
	u.Email = email
	err := runUserValidatorFuncs(&u, uv.emailNormalize)
	if err != nil {
		return nil, err
	}

	return uv.UserDB.FindByEmail(email)
}

func (uv *userValidator) Create(user *User) error {
	err := runUserValidatorFuncs(user,
		uv.usernameRequire,
		uv.usernameNormalize,
		uv.usernameIsAvailable,
		uv.emailRequire,
		uv.emailNormalize,
		uv.emailCheckFormat,
		uv.emailIsAvailable,
		uv.passwordRequire,
	)
	if err != nil {
		return err
	}

	return uv.UserDB.Create(user)
}

func (uv *userValidator) Update(user *User) error {
	// TODO

	return uv.UserDB.Update(user)
}

func (uv *userValidator) Delete(id uint) error {
	user := User{Model: gorm.Model{ID: id}}
	err := runUserValidatorFuncs(&user, uv.idGreaterThanZero)
	if err != nil {
		return err
	}

	return uv.UserDB.Delete(id)
}

type userGorm struct {
	db *gorm.DB
}

func newUserGorm(db *gorm.DB) *userGorm {
	return &userGorm{
		db: db,
	}
}

func (ug *userGorm) FindByUsername(username string) (*User, error) {
	u := User{}
	err := ug.db.Where("username = ?", username).First(&u).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

func (ug *userGorm) FindByEmail(email string) (*User, error) {
	u := User{}
	err := ug.db.Where("email = ?", email).First(&u).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errRecordNotFound
	} else if err != nil {
		return nil, err
	}
	return &u, nil
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

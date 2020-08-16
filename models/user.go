package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres driver
	"golang.org/x/crypto/bcrypt"

	"quacker/hash"
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
	About             string
}

// UserDB is an interface for interacting with user data in the database
type UserDB interface {
	FindByUsername(username string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByRememberToken(token string) (*User, error)

	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

// UserService is an interface for interacting with user model
type UserService interface {
	UserDB
	Authenticate(login, password string) (*User, error)
}

type userService struct {
	UserDB
	passwordPepper string
}

// NewUserService creates UserService instance
func NewUserService(db *gorm.DB, passwordPepper, hmacKey string) UserService {
	ug := newUserGorm(db)
	uv := newUserValidator(ug, passwordPepper, hmacKey)

	return &userService{
		UserDB:         uv,
		passwordPepper: passwordPepper,
	}
}

func (us *userService) Authenticate(login, password string) (*User, error) {
	user, err := us.FindByEmail(login)
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password+us.passwordPepper))
		if err != nil {
			return nil, ErrCredentialsInvalid
		}
		return user, nil
	}

	user, err = us.FindByUsername(login)
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password+us.passwordPepper))
		if err != nil {
			return nil, ErrCredentialsInvalid
		}
		return user, nil
	}

	return nil, ErrCredentialsInvalid
}

type userValidator struct {
	UserDB
	EmailRegex     *regexp.Regexp
	UsernameRegex  *regexp.Regexp
	PasswordPepper string
	Hmac           hash.Hmac
}

func newUserValidator(u UserDB, passwordPepper, hmacKey string) *userValidator {
	return &userValidator{
		UserDB:         u,
		EmailRegex:     regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
		UsernameRegex:  regexp.MustCompile(`^[a-zA-Z0-9_-]+`),
		PasswordPepper: passwordPepper,
		Hmac:           hash.NewHmac(hmacKey),
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

func (uv *userValidator) FindByRememberToken(token string) (*User, error) {
	u := User{}
	u.RememberToken = token
	err := runUserValidatorFuncs(&u,
		uv.rememberTokenHashCreate,
	)
	if err != nil {
		return nil, err
	}

	return uv.UserDB.FindByRememberToken(u.RememberTokenHash)
}

func (uv *userValidator) Create(user *User) error {
	err := runUserValidatorFuncs(user,
		uv.usernameRequire,
		uv.usernameNormalize,
		uv.usernameCheckFormat,
		uv.usernameIsAvailable,
		uv.emailRequire,
		uv.emailNormalize,
		uv.emailCheckFormat,
		uv.emailIsAvailable,
		uv.passwordRequire,
		uv.passwordHashCreate,
		uv.passwordHashRequire,
		uv.rememberTokenCreate,
		uv.rememberTokenHashCreate,
		uv.rememberTokenHashRequire,
	)
	if err != nil {
		return err
	}

	return uv.UserDB.Create(user)
}

func (uv *userValidator) Update(user *User) error {
	err := runUserValidatorFuncs(user,
		uv.usernameRequire,
		uv.usernameNormalize,
		uv.usernameCheckFormat,
		uv.usernameIsAvailable,
		uv.emailRequire,
		uv.emailNormalize,
		uv.emailCheckFormat,
		uv.emailIsAvailable,
		uv.passwordHashCreate,
		uv.passwordHashRequire,
		uv.rememberTokenHashCreate,
		uv.rememberTokenHashRequire,
	)
	if err != nil {
		return err
	}

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
		return nil, ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

func (ug *userGorm) FindByEmail(email string) (*User, error) {
	u := User{}
	err := ug.db.Where("email = ?", email).First(&u).Error

	if err == gorm.ErrRecordNotFound {
		return nil, ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}

func (ug *userGorm) FindByRememberToken(token string) (*User, error) {
	u := User{}
	err := ug.db.Where("remember_token_hash = ?", token).First(&u).Error

	if err == gorm.ErrRecordNotFound {
		return nil, ErrRecordNotFound
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

package models

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userValidatorFunc func(*User) error

func runUserValidatorFuncs(user *User, fns ...userValidatorFunc) error {
	for _, f := range fns {
		err := f(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uv *userValidator) idGreaterThanZero(user *User) error {
	if user.ID <= 0 {
		return errInvalidID
	}
	return nil
}

func (uv *userValidator) usernameUnique(user *User) error {
	// TODO
	return nil
}

func (uv *userValidator) emailNormalize(user *User) error {
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	return nil
}

func (uv *userValidator) emailRequire(user *User) error {
	if user.Email == "" {
		return errEmailRequired
	}
	return nil
}

func (uv *userValidator) emailCheckFormat(user *User) error {
	// TODO
	return nil
}

func (uv *userValidator) emailIsUnique(user *User) error {
	// TODO
	return nil
}

func (uv *userValidator) passwordRequire(user *User) error {
	if user.Password == "" {
		return errPasswordRequired
	}
	return nil
}

func (uv *userValidator) passwordHashRequire(user *User) error {
	if user.PasswordHash == "" {
		return errPasswordHashRequired
	}
	return nil
}

func passwordEncrypt(user *User) error {
	if user.Password == "" {
		return errPasswordRequired
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	return nil
}

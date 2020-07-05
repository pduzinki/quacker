package models

import (
	"strings"
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

	return nil
}

func (uv *userValidator) passwordRequire(user *User) error {
	if user.Password == "" {
		return errPasswordRequired
	}
	return nil
}

func (uv *userValidator) passwordHashRequire(user *User) error {
	return nil
}

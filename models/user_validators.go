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

func (uv *userValidator) usernameNormalize(user *User) error {
	user.Username = strings.ToLower(user.Username)
	user.Username = strings.TrimSpace(user.Username)
	return nil
}

func (uv *userValidator) usernameIsAvailable(user *User) error {
	existingUser, err := uv.FindByUsername(user.Username)
	if err == errRecordNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	if existingUser.ID != user.ID {
		return errUsernameTaken
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
	// TODO
	return nil
}

func (uv *userValidator) emailIsAvailable(user *User) error {
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

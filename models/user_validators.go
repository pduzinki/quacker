package models

import (
	"strings"

	"golang.org/x/crypto/bcrypt"

	"quacker/token"
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
		return ErrInvalidID
	}
	return nil
}

func (uv *userValidator) usernameRequire(user *User) error {
	if user.Username == "" {
		return ErrUsernameRequired
	}
	return nil
}

func (uv *userValidator) usernameNormalize(user *User) error {
	user.Username = strings.ToLower(user.Username)
	user.Username = strings.TrimSpace(user.Username)
	return nil
}

func (uv *userValidator) usernameCheckFormat(user *User) error {
	if !uv.UsernameRegex.MatchString(user.Username) {
		return ErrUsernameInvalidFormat
	}
	return nil
}

func (uv *userValidator) usernameIsAvailable(user *User) error {
	existingUser, err := uv.FindByUsername(user.Username)
	if err == ErrRecordNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	if existingUser.ID != user.ID {
		return ErrUsernameTaken
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
		return ErrEmailRequired
	}
	return nil
}

func (uv *userValidator) emailCheckFormat(user *User) error {
	if !uv.EmailRegex.MatchString(user.Email) {
		return ErrEmailInvalidFormat
	}
	return nil
}

func (uv *userValidator) emailIsAvailable(user *User) error {
	existingUser, err := uv.FindByEmail(user.Email)
	if err == ErrRecordNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	if existingUser.ID != user.ID {
		return ErrEmailTaken
	}

	return nil
}

func (uv *userValidator) passwordRequire(user *User) error {
	if user.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}

func (uv *userValidator) passwordHashCreate(user *User) error {
	if user.Password == "" {
		return nil // no new password given, do nothing
	}

	passwordBytes := user.Password + uv.PasswordPepper
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passwordBytes), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	return nil
}

func (uv *userValidator) passwordHashRequire(user *User) error {
	if user.PasswordHash == "" {
		return ErrPasswordHashRequired
	}
	return nil
}

func (uv *userValidator) rememberTokenCreate(user *User) error {
	token, err := token.GenerateRememberToken()
	if err != nil {
		return err
	}

	user.RememberToken = token
	return nil
}

func (uv *userValidator) rememberTokenHashCreate(user *User) error {
	if user.RememberToken == "" {
		return nil // no new token created, do nothing
	}

	tokenHash := uv.Hmac.Hash(user.RememberToken)
	user.RememberTokenHash = tokenHash
	return nil
}

func (uv *userValidator) rememberTokenHashRequire(user *User) error {
	if user.RememberTokenHash == "" {
		return ErrRememberTokenHashRequired
	}
	return nil
}

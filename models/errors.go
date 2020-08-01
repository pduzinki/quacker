package models

import "errors"

var (
	ErrInvalidID                 = errors.New("model: Invalid ID")
	ErrEmailRequired             = errors.New("model: Email required")
	ErrEmailTaken                = errors.New("model: Email taken")
	ErrEmailInvalidFormat        = errors.New("model: Email invalid format")
	ErrPasswordRequired          = errors.New("model: Password required")
	ErrPasswordHashRequired      = errors.New("model: Password hash required")
	ErrRecordNotFound            = errors.New("model: Record not found")
	ErrUsernameTaken             = errors.New("model: Username taken")
	ErrUsernameRequired          = errors.New("model: Username required")
	ErrUsernameInvalidFormat     = errors.New("model: Username invalid format")
	ErrRememberTokenHashRequired = errors.New("model: Remember token hash required")
)

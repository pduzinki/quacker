package models

import "errors"

var (
	errInvalidID            = errors.New("model: Invalid ID")
	errEmailRequired        = errors.New("model: Email required")
	errEmailInvalidFormat   = errors.New("model: Email invalid format")
	errPasswordRequired     = errors.New("model: Password required")
	errPasswordHashRequired = errors.New("model: Password hash required")
	errRecordNotFound       = errors.New("model: Record not found")
	errUsernameTaken        = errors.New("model: Username taken")
	errUsernameRequired     = errors.New("model: Username required")
)

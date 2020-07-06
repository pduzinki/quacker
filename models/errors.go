package models

import "errors"

var (
	errInvalidID            = errors.New("model: Invalid ID")
	errEmailRequired        = errors.New("model: Email required")
	errPasswordRequired     = errors.New("model: Password required")
	errPasswordHashRequired = errors.New("model: Password hash required")
	errRecordNotFound       = errors.New("model: Record not found")
	errUsernameTaken        = errors.New("model: Username taken")
)

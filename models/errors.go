package models

import "errors"

var (
	errInvalidID        = errors.New("model: Invalid ID")
	errEmailRequired    = errors.New("model: Email required")
	errPasswordRequired = errors.New("model: Password required")
)

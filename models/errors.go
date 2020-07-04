package models

import "errors"

var (
	errInvalidID     = errors.New("model: Invalid ID")
	errEmailRequired = errors.New("model: Email required")
)

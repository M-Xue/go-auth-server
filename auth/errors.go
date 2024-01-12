package auth

import (
	"errors"
)

var (
	ErrInvalidToken = errors.New("authentication token is invalid")
	ErrExpiredToken = errors.New("authentication token has expired")
)

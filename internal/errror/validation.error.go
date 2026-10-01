package apperror

import (
	"errors"
)

var (
	ErrEmptyField         = errors.New("Email or password can't be empty")
	ErrInvalidEmailFormat = errors.New("Invalid Email Format")
	ErrAlreadyExist       = errors.New("Email already exist")
	ErrWrongEmailPass     = errors.New("Wrong email or password")
	ErrFileFormat         = errors.New("image format must be jpg/jpeg/png")
	ErrFileSize           = errors.New("image max size 2mb")
	ErrWrongPass          = errors.New("password lama salah")
)

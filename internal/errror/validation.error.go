package apperror

import (
	"errors"
)

var (
	ErrEmptyField         = errors.New("Email or password can't be empty")
	ErrInvalidEmailFormat = errors.New("Invalid Email Format")
	ErrAlreadyExist       = errors.New("Email already exist")
	ErrWrongEmailPass     = errors.New("Wrong email or password")
	ErrFileFormat         = errors.New("format foto harus jpg/jpeg/png")
	ErrFileSize           = errors.New("ukuran foto maksimal 2mb")
)

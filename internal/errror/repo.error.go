package apperror

import "errors"

var (
	ErrNoRowsAffected = errors.New("no row affected")
)

package error

import "errors"

var (
	// ErrNoRow example
	ErrNoRow = errors.New("no rows in result set")
)
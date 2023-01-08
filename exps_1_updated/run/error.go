package run

import "errors"

var (
	ErrNOutOfRange = errors.New("n: out of range")
	ErrPOutOfRange = errors.New("p: out of range")
	ErrInvalidMode = errors.New("invalid mode")
)

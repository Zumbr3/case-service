package errors

import "errors"

var (
	ErrNotImplemented    = errors.New("not implemented")
	ErrCaseNotFound      = errors.New("case not found")
	ErrCaseAlreadyExists = errors.New("case already exists")
)

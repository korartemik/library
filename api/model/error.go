package model

type NotFoundError interface {
	error
	NotFoundError()
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(NotFoundError)
	return ok
}

func NotFound(err error) NotFoundError {
	return &notFoundError{err}
}

type notFoundError struct {
	cause error
}

var _ NotFoundError = (*notFoundError)(nil)

func (e *notFoundError) Cause() error {
	return e.cause
}

func (e *notFoundError) Error() string {
	return e.cause.Error()
}

func (e *notFoundError) NotFoundError() {}

package httperr

import "net/http"

type ErrorHTTP interface {
	Code() int
	Reason() string
	Error() string
}

type httperror struct {
	statusCode int
	reason     string
	err        error
}

func (e httperror) Code() int      { return e.statusCode }
func (e httperror) Reason() string { return e.reason }
func (e httperror) Error() string  { return e.err.Error() }

func InternalServerError(reason string, err error) ErrorHTTP {
	return &httperror{
		statusCode: http.StatusInternalServerError,
		reason:     reason,
		err:        err,
	}
}

func BadRequest(reason string, err error) ErrorHTTP {
	return &httperror{
		statusCode: http.StatusBadRequest,
		reason:     reason,
		err:        err,
	}
}

func Unauthorized(reason string, err error) ErrorHTTP {
	return &httperror{
		statusCode: http.StatusUnauthorized,
		reason:     reason,
		err:        err,
	}
}

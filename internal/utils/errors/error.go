package errors

import "net/http"

type ErrorString struct {
	code    int
	status  string
	message string
}

func (e ErrorString) Code() int {
	return e.code
}

func (e ErrorString) Error() string {
	return e.message
}

func (e ErrorString) Message() string {
	return e.message
}

func BadRequest(msg string) error {
	return &ErrorString{
		code:    http.StatusBadRequest,
		status:  http.StatusText(http.StatusBadRequest),
		message: msg,
	}
}

func NotFound(msg string) error {
	return &ErrorString{
		code:    http.StatusNotFound,
		status:  http.StatusText(http.StatusNotFound),
		message: msg,
	}
}

func Conflict(msg string) error {
	return &ErrorString{
		code:    http.StatusConflict,
		status:  http.StatusText(http.StatusConflict),
		message: msg,
	}
}

func InternalServerError(msg string) error {
	return &ErrorString{
		code:    http.StatusInternalServerError,
		status:  http.StatusText(http.StatusConflict),
		message: msg,
	}
}

func UnauthorizedError(msg string) error {
	return &ErrorString{
		code:    http.StatusUnauthorized,
		status:  http.StatusText(http.StatusUnauthorized),
		message: msg,
	}
}

func ForbiddenError(msg string) error {
	return &ErrorString{
		code:    http.StatusForbidden,
		status:  http.StatusText(http.StatusForbidden),
		message: msg,
	}
}

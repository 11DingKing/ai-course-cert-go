package apperr

import (
	"errors"
	"net/http"
)

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		switch e.Code {
		case Invalid:
			return 400
		case Unauthorized:
			return 401
		case Forbidden:
			return 403
		case NotFound:
			return 404
		case Conflict:
			return 409
		case Deadline:
			return 408
		}
	}
	return http.StatusInternalServerError
}
func Message(err error) string {
	var e *Error
	if errors.As(err, &e) {
		return e.Message
	}
	return "internal error"
}
func IsCode(err error, c Code) bool          { var e *Error; return errors.As(err, &e) && e.Code == c }
func Wrap(c Code, m string, err error) error { return E(c, m, err) }

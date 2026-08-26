package apperr

import "fmt"

type Code string

const (
	Invalid      Code = "invalid"
	NotFound     Code = "not_found"
	Conflict     Code = "conflict"
	Forbidden    Code = "forbidden"
	Unauthorized Code = "unauthorized"
	Deadline     Code = "deadline"
	Internal     Code = "internal"
)

type Error struct {
	Code    Code
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
func (e *Error) Unwrap() error            { return e.Err }
func E(c Code, m string, err error) error { return &Error{Code: c, Message: m, Err: err} }

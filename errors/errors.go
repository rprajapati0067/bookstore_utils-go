package errors

import "net/http"

type Errors interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type error struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e error) Message() string {
	panic("implement me")
}

func (e error) Status() int {
	panic("implement me")
}

func (e error) Error() string {
	panic("implement me")
}

func (e error) Causes() []interface{} {
	panic("implement me")
}

func NewNotFoundError(message string) Errors {
	var result Errors
	result = error{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}

	return result
}

func NewInternalServerError(message string, err string) Errors {
	result := error{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}

	return result
}

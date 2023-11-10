package util

import (
	"net/http"
)

type Error struct {
	Message string
	Code    int
	Err     error
	Data    interface{}
}

// setter
func (e *Error) Set(code int, err error, message string) {
	e.Code = code
	e.Message = message
	e.Err = err
}

func (e *Error) SetError(err error) {
	if err != nil {
		e.Err = err
		e.Message = err.Error()
	}
}

func (e *Error) StatusBadRequest(err error, message string) {
	e.Code = http.StatusBadRequest
	e.Message = message
	e.Err = err
}

func (e *Error) StatusUnprocessableEntity(err error, message string) {
	e.Code = http.StatusUnprocessableEntity
	e.Message = message
	e.Err = err
}

func (e *Error) StatusNotFound(err error, message string) {
	e.Code = http.StatusNotFound
	e.Message = message
	e.Err = err
}

func (e *Error) StatusUnauthorized(err error, message string) {
	e.Code = http.StatusUnauthorized
	e.Message = message
	e.Err = err
}

func (e *Error) StatusInternalServerError(err error, message string) {
	e.Code = http.StatusInternalServerError
	e.Message = message
	e.Err = err
}

func (e *Error) StatusMethodNotAllowed(err error, message string) {
	e.Code = http.StatusMethodNotAllowed
	e.Message = message
	e.Err = err
}

// getter
func (e Error) GetError() error {
	return e.Err
}

func (e Error) GetCode() int {
	return e.Code
}

func (e Error) GetMessage() string {
	return e.Message
}

func (e Error) GetData() interface{} {
	return e.Data
}

// nil
func (e Error) NotNil() bool {
	if e.Code != 0 || e.Message != "" || e.Err != nil {
		return true
	}
	return false
}

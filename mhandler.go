package mhandler

import (
	"fmt"
	"log"
	"net/http"
)

type Error struct {
	Error   error
	Message string
	Code    int
}

type Handler func(http.ResponseWriter, *http.Request) *Error

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("%+v: %s", e.Error, e.Message)
		jsonErrorString := fmt.Sprintf(`{"errors":[{"msg":"%s","code":"%d"}]}`, e.Message, e.Code)
		http.Error(w, jsonErrorString, e.Code)
	}
}

func ErrorInternal(err error) *Error {
	return &Error{
		Error:   err,
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
	}
}

func ErrorBadRequest(err error) *Error {
	return &Error{
		Error:   err,
		Message: err.Error(),
		Code:    http.StatusBadRequest,
	}
}

func ErrorNotFound(err error) *Error {
	return &Error{
		Error:   err,
		Message: err.Error(),
		Code:    http.StatusNotFound,
	}
}

func ErrorForbidden(err error) *Error {
	return &Error{
		Error:   err,
		Message: err.Error(),
		Code:    http.StatusForbidden,
	}
}

package mhandler

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	glog "google.golang.org/appengine/log"
)

type Error struct {
	Error   error
	Message string
	Code    int
}

type Handler func(http.ResponseWriter, *http.Request) *Error

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		c := appengine.NewContext(r)
		glog.Errorf(c, "%+v: %s", e.Error, e.Message)
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

package handlers

import (
	"net/http"
	"strings"
)

// Error represents a HTTP error
type Error struct {
	HTTPCode int    `json:"-"`
	Error    string `json:"-"`
}

func newError(HTTPCode int, msg string) func(...string) (int, *Error) {
	return func(msgs ...string) (int, *Error) {
		if msgs != nil && len(msgs) != 0 {
			msg = strings.Join(msgs, " ")
		}
		return HTTPCode, &Error{HTTPCode, msg}
	}
}

// All the newError constants would be here
var (
	ErrNotFound                    = newError(http.StatusNotFound, "Not Found")
	ErrBadRequestParametersMissing = newError(http.StatusBadRequest, "Mandatory Parameter missing")
	ErrBadRequestInvalidParameter  = newError(http.StatusBadRequest, "Invalid parameter")
	ErrBadRequestInvalidBody       = newError(http.StatusBadRequest, "Invalid Body")
	ErrInternalServerError         = newError(http.StatusInternalServerError, "Internal Server Error")
	ErrMethodNotDefined            = newError(http.StatusMethodNotAllowed, "Method not implemented")
	ErrResourceConflict            = newError(http.StatusConflict, "Resource Conflict")
	ErrInvalidOperation            = newError(http.StatusUnavailableForLegalReasons, "Unavailable for leagal reasons")
)

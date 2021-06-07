package cotter

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrUnauthorized represents token validation errors.
const ErrUnauthorized e = "unauthorized"

// ErrorHandler function.
type ErrorHandler func(w http.ResponseWriter, r *http.Request, e error)

type e string

func (e) Error() string { return "" }

func unauthorized(message string) error {
	return fmt.Errorf("%s%w", message, ErrUnauthorized)
}

func defaultErrorHandler(w http.ResponseWriter, r *http.Request, e error) {
	if e == nil {
		return
	}

	if errors.Is(e, ErrUnauthorized) {
		http.Error(w, e.Error(), http.StatusUnauthorized)
		return
	}

	http.Error(w, e.Error(), http.StatusInternalServerError)
}

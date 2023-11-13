package httpErrors

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound       = errors.New("Not Found")
	ErrValidate       = errors.New("Validation error")
	ErrUnexpected     = errors.New("Unexpected error")
	ErrTimeout        = errors.New("Timeout exceeded")
	ErrEmailExist     = errors.New("User already exists")
	ErrUnauthorized   = errors.New("Unauthorized")
	ErrBadRequest     = errors.New("Bad request")
	ErrInternalServer = errors.New("Internal Server Error")
)

type Error struct {
	ErrStatus  int    `json:"status,omitempty"`
	ErrMessage string `json:"message,omitempty"`
	ErrCauses  any    `json:"-"`
}

func (s Error) Error() string {
	return fmt.Sprintf("status: %d, message: %s", s.ErrStatus, s.ErrMessage)
}

func NewNotFound(causes any) Error {
	return Error{http.StatusNotFound, ErrNotFound.Error(), causes}
}

func NewBadRequest(causes any) Error {
	return Error{http.StatusBadRequest, ErrBadRequest.Error(), causes}
}

func NewUnauthorized(causes any) Error {
	return Error{http.StatusUnauthorized, ErrUnauthorized.Error(), causes}
}

func ParseError(err error) Error {
	if errors, ok := err.(validator.ValidationErrors); ok {
		mess := fmt.Sprintf("%s %s", strings.ToLower(errors[0].Field()), errors[0].Tag())
		return Error{http.StatusBadRequest, mess, nil}
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return Error{http.StatusNotFound, ErrNotFound.Error(), nil}
	case errors.Is(err, strconv.ErrSyntax):
		return Error{http.StatusBadRequest, ErrValidate.Error(), nil}
	case errors.Is(err, context.DeadlineExceeded):
		return Error{http.StatusRequestTimeout, ErrTimeout.Error(), nil}
	default:
		if err, ok := err.(Error); ok {
			return err
		}
		return Error{http.StatusInternalServerError, ErrUnexpected.Error(), nil}
	}
}

func RequestError(err error) (int, Error) {
	e := ParseError(err)
	return e.ErrStatus, e
}

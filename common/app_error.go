package common

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewFullErrorResponse(statusCode int, rootError error, message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootError,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(rootError error, message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    rootError,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorizedErrorResponse(rootError error, message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    rootError,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func (e *AppError) RootError() error {
	var err *AppError
	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewCustomErrorResponse(root error, message string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, message, root.Error(), key)
	}
	return NewErrorResponse(root, message, message, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong in DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "INVALID_REQUEST_ERROR")
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(err, fmt.Sprintf("cannot create entity %s", entity), fmt.Sprintf("CANNOT_CREATE_ENTITY_%s", entity))
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(err, fmt.Sprintf("cannot get entity %s", entity), fmt.Sprintf("CANNOT_GET_ENTITY_%s", entity))
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(err, fmt.Sprintf("cannot update entity %s", entity), fmt.Sprintf("CANNOT_UPDATE_ENTITY_%s", entity))
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(err, fmt.Sprintf("cannot delete entity %s", entity), fmt.Sprintf("CANNOT_DELETE_ENTITY_%s", entity))
}

var RecordNotFoundError = errors.New("record not found")

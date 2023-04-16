package errors

import "fmt"

var (
	ErrAuthCredentials = newError(en[E_AUTH_CREDENTIALS], E_AUTH_CREDENTIALS)
	ErrUserNotFound    = newError(en[E_USER_NOT_FOUND], E_USER_NOT_FOUND)
	ErrInternalServer  = newError(en[E_INTERNAL_SERVER], E_INTERNAL_SERVER)
	ErrAuthJwt         = newInternalServerError("jwt", en[E_AUTH_JWT], E_AUTH_JWT)
)

// Reference: "io/fs".PathError
// https://cs.opensource.google/go/go/+/master:src/io/fs/fs.go;l=244?q=PathError&sq=&ss=go%2Fgo

type (
	ErrCode  string
	AppError struct {
		Message string  `json:"message"`
		Code    ErrCode `json:"code"`
		Op      string  `json:"operation"`
		Err     error   `json:"meta"`
	}
)

func newError(message string, code ErrCode) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func (e *AppError) Operation(op string) *AppError {
	e.Op = op
	return e
}

func (e *AppError) Wrap(err error) *AppError {
	e.Err = err
	return e
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s (%s)", e.Message, e.Code)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// WrapDbError is used to wrap errors coming out of database operations
func WrapDbError(err error) *AppError {
	return newError(en[E_DB_OPERATION], E_DB_OPERATION).
		Operation("database").
		Wrap(err)
}

func newInternalServerError(operation, message string, code ErrCode) *AppError {
	return newError(message, code).
		Operation(operation).
		Wrap(ErrInternalServer)
}

// WrapInternalServerError is used to wrap server errors coming out of internal ops
func WrapInternalServerError(operation string, err error) *AppError {
	return newError(en[E_INTERNAL_SERVER], E_INTERNAL_SERVER).
		Operation(operation).
		Wrap(fmt.Errorf("%w %w", ErrInternalServer, err))
}

const (
	E_AUTH_CREDENTIALS ErrCode = "E_AUTH_CREDENTIALS"
	E_USER_NOT_FOUND           = "E_USER_NOT_FOUND"
	E_DB_OPERATION             = "E_DB_ERR"
	E_INTERNAL_SERVER          = "E_INTERNAL_SERVER"
	E_AUTH_JWT                 = "E_AUTH_JWT"
)

type translation map[ErrCode]string

var en translation = map[ErrCode]string{
	E_AUTH_CREDENTIALS: "wrong username or password",
	E_USER_NOT_FOUND:   "user not found",
	E_DB_OPERATION:     "database operation issue",
	E_INTERNAL_SERVER:  "internal server error",
	E_AUTH_JWT:         "unable to issue jwt",
}

// Translations table
var _ map[string]translation = map[string]translation{
	"en": en,
	"id": nil,
}

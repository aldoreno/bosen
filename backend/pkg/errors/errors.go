package errors

import "fmt"

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

const (
	E_AUTH_CREDENTIALS  ErrCode = "E_AUTH_CREDENTIALS"
	E_ACCOUNT_NOT_FOUND         = "E_AUTH_ACCOUNT"
	E_DB_OPERATION              = "E_AUTH_DB_ERR"
)

type translation map[ErrCode]string

var en translation = map[ErrCode]string{
	E_AUTH_CREDENTIALS:  "wrong username or password",
	E_ACCOUNT_NOT_FOUND: "account not found",
	E_DB_OPERATION:      "database operation issue",
}

// Translations table
var _ map[string]translation = map[string]translation{
	"en": en,
	"id": nil,
}

var (
	ErrWrongUsernameOrPassword = newError(en[E_AUTH_CREDENTIALS], E_AUTH_CREDENTIALS)
	ErrAccountNotFound         = newError(en[E_ACCOUNT_NOT_FOUND], E_ACCOUNT_NOT_FOUND)
)

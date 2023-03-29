package login

import "fmt"

// Reference: "io/fs".PathError
// https://cs.opensource.google/go/go/+/master:src/io/fs/fs.go;l=244?q=PathError&sq=&ss=go%2Fgo

type (
	ErrorCode string
	AuthError struct {
		Message string    `json:"message"`
		Code    ErrorCode `json:"code"`
	}
)

func newError(message string, code ErrorCode) *AuthError {
	return &AuthError{
		Message: message,
		Code:    code,
	}
}

func (e *AuthError) Error() string { return fmt.Sprintf("%s (%s)", e.Message, e.Code) }

const (
	wrongUsernameOrPassword ErrorCode = "E_AUTH_CREDENTIALS"
	accountNotFound                   = "E_AUTH_ACCOUNT"
)

type translation map[ErrorCode]string

var en translation = map[ErrorCode]string{
	wrongUsernameOrPassword: "wrong username or password",
	accountNotFound:         "account not found",
}

// Translations table
var _ map[string]translation = map[string]translation{
	"en": en,
	"id": nil,
}

var (
	ErrWrongUsernameOrPassword = newError(en[wrongUsernameOrPassword], wrongUsernameOrPassword)
	ErrAccountNotFound         = newError(en[accountNotFound], accountNotFound)
)

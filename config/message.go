package config

const (
	OK                  = 0
	InternalServerError = 10001
	ErrBind             = 10002
	UnverifiedError_1   = 10003
	UnverifiedError_2   = 10004
	ErrUserNotFound     = 20102
)

var codeText = map[int]string{
	OK:                  "OK",
	InternalServerError: "Internal server error",
	ErrBind:             "Error occurred while binding the request body to the struct.",
	UnverifiedError_1:   "JWT is not exist.",
	UnverifiedError_2:   "JWT is invalid.",
	ErrUserNotFound:     "The user was not found.",
}

func CodeText(code int) string {
	return codeText[code]
}

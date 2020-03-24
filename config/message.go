package config

const (
	OK = 0

	InternalServerError = 10001
	ErrBind             = 10002
	CreateTokenError    = 10010

	UnverifiedError_0 = 20002
	UnverifiedError_1 = 20003
	UnverifiedError_2 = 20004
	ErrUserNotFound   = 20102
)

var codeText = map[int]string{
	OK:                  "OK",
	InternalServerError: "Internal server error",
	ErrBind:             "Error occurred while binding the request body to the struct.",
	UnverifiedError_0:   "Email or password is wrong.",
	UnverifiedError_1:   "JWT is not exist.",
	UnverifiedError_2:   "JWT is invalid.",
	CreateTokenError:    "An internal server error occurred when JTW token was created.",
	ErrUserNotFound:     "The user was not found.",
}

func CodeText(code int) string {
	return codeText[code]
}

package config

const (
	OK = 0

	InternalServerError = 10001
	ErrBind             = 10002
	CreateTokenError    = 10010

	UnverifiedError = 20002
	UnverifiedError_JWTNonexist = 20003
	UnverifiedError_JWTInvalid = 20004
	UnverifiedError_JWTExpired = 20005
	ErrUserNotFound   = 20102
)

var codeText = map[int]string{
	OK:                  "OK",
	InternalServerError: "Internal server error",
	ErrBind:             "Error occurred while binding the request body to the struct.",
	UnverifiedError:   "Email or password is wrong.",
	UnverifiedError_JWTNonexist:   "JWT is not exist.",
	UnverifiedError_JWTInvalid:   "JWT is invalid.",
	UnverifiedError_JWTExpired:   "JWT is Expired.",
	CreateTokenError:    "An internal server error occurred when JTW token was created.",
	ErrUserNotFound:     "The user was not found.",
}

func CodeText(code int) string {
	return codeText[code]
}

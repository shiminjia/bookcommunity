package config

const (
	OK          = 0
	EmailVerify = 00001

	InternalServerError = 10001
	ErrBind             = 10002
	ValidError          = 10003
	UniqueErr           = 10004
	CreateTokenError    = 10010
	DBErr               = 10020
	EmailUniqueErr      = 10021
	UsernameUniqueErr   = 10022

	EmailVerifyErr = 10030

	UnverifiedError             = 20001
	UnverifiedError_JWTNonexist = 20002
	UnverifiedError_JWTNoBearer = 20003
	UnverifiedError_JWTInvalid  = 20004
	UnverifiedError_JWTExpired  = 20005
	UnverifiedError_JWTScopeErr = 20006

	ErrUserNotFound = 20102
	LOGINFailure    = 20103
	ErrGender       = 20104

	Forbidden    = 20201
)

var codeText = map[int]string{
	OK:          "OK",
	EmailVerify: "verification successful.",

	InternalServerError: "Internal server error",
	ErrBind:             "Error occurred while binding the request body to the struct.",
	ValidError:          "Please check your input value. Some values do not meet requirements",
	UniqueErr:           "Email or username is not unique",
	CreateTokenError:    "An internal server error occurred when JTW token was created.",
	DBErr:               "DB access error. Please contact the administrator",
	EmailUniqueErr:      "Email is not unique.",
	UsernameUniqueErr:   "Username is not unique.",

	EmailVerifyErr: "Email Verification is completed.",

	UnverifiedError:             "Email or password is wrong.",
	UnverifiedError_JWTNonexist: "JWT is not exist.",
	UnverifiedError_JWTNoBearer: "No Bearer at the head of JWT. Please check it add Bearer like Bearer XXXXXXXXXXX.",
	UnverifiedError_JWTInvalid:  "JWT is invalid.",
	UnverifiedError_JWTExpired:  "JWT is Expired.",
	UnverifiedError_JWTScopeErr: "JWT scope is invalid.",

	ErrUserNotFound: "The user was not found.",
	LOGINFailure:    "Email or Password is wrong.",
	ErrGender:      "Gender is invalid.",

	Forbidden:   "Permission denied.",
}

func CodeText(code int) string {
	return codeText[code]
}

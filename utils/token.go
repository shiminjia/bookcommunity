package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Context struct {
	Id       string
	Email    string
	Username string
}

func CreateToken(ctx *Context)(tokenString string, err error){
	secret := os.Getenv("JWT_SECRET")

	//claim a new token by NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       ctx.Id,
		"username": ctx.Username,
		"nbf":      time.Now().Unix(),    //creation time
		"iat":      time.Now().Unix(),    //
	})

	//sign the token
	tokenString, err = token.SignedString([]byte(secret))
	return
}

func VerifyToken(token string) bool {

	return true

}



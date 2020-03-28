package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/shiminjia/bookcommunity/config"
	"os"
	"strconv"
	"time"
)

type Context struct {
	Iss      string
	Exp      int64
	Sub      string
	Nbf      int64
	Iat      int64
	Jti      string
	Id       string
	Username string
}

func CreateToken(ctx *Context) (tokenString string, err error) {
	secret := config.JWT_SECRET
	iss := config.ISSUER
	eft := config.JWT_EFFECTIVE_TIME

	//change eft to eft64. eft means effective time of jwt.
	eft64, err := strconv.ParseInt(eft, 10, 64)
	if err != nil {
		eft64 = 24 * 60 * 60                 //one day
	}

	//claim a new token by NewWithClaims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      iss,                            //issuer
		"exp":      time.Now().Unix() + eft64,      //expiration time
		"sub":      "",                             //subject (not decided now)
		"nbf":      time.Now().Unix(),              //Not Before
		"iat":      time.Now().Unix(),              //Issued At
		"jti":      uuid.New(),                     //JWT ID
		"id":       ctx.Id,                         //user id
		"username": ctx.Username,                   //user name
	})

	//sign the token
	tokenString, err = token.SignedString([]byte(secret))
	return
}

func ParseToken(tokenString string) (*Context, error) {
	ctx := &Context{}

	//get JTW_SECRET
	secret := os.Getenv("JWT_SECRET")

	//prase tokenString is vaild or not.
	//secretFunc(secret) means
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil{
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Iss = claims["iss"].(string)
		ctx.Exp = int64(claims["exp"].(float64))
		ctx.Sub = claims["sub"].(string)
		ctx.Nbf = int64(claims["nbf"].(float64))
		ctx.Iat = int64(claims["iat"].(float64))
		ctx.Jti = claims["jti"].(string)
		ctx.Id = claims["id"].(string)
		ctx.Username = claims["username"].(string)
		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}



package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/db"
	"github.com/shiminjia/bookcommunity/middleware"
	"github.com/shiminjia/bookcommunity/utils"
	"net/http"
	"strconv"
	"time"
)

func EmailVerification(c *gin.Context){
	//get token from query
	token := c.Query("token")
	if token == "" {
		middleware.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTNonexist)
		return
	}

	//Parse and verify token
	ctx, err := utils.ParseToken(token)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTInvalid)
		return
	}

	//check scope
	if ctx.Scope != "email" {
		middleware.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTScopeErr)
		return
	}

	//Check if token is within the validity period
	if time.Now().Unix() > ctx.Exp {
		middleware.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTExpired)
		return
	}

	IdInt, err := strconv.Atoi(ctx.Id)
	IdUnit := uint(IdInt)

	//create a user model to update status
	user := db.User{
		Id:           db.Id{IdUnit},
		Status:       1,
	}

	//Check if status is 0(unverified) or not
	res, err := user.GetUserById()
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}
	if res.Status != 0 {
		middleware.ErrorResponse(c, http.StatusForbidden, config.EmailVerifyErr)
		return
	}


	_, err = user.UpdateStatus()
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}

	middleware.ErrorResponse(c, http.StatusOK, config.EmailVerify)

}

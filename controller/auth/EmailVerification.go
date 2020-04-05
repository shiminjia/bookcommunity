package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/db"
	"github.com/shiminjia/bookcommunity/utils"
	"net/http"
	"time"
)

func EmailVerification(c *gin.Context) {
	//get token from query
	token := c.Query("token")
	if token == "" {
		utils.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTNonexist)
		return
	}

	//Parse and verify token
	ctx, err := utils.ParseToken(token)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTInvalid)
		return
	}

	//check scope
	if ctx.Scope != "email" {
		utils.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTScopeErr)
		return
	}

	//Check if token is within the validity period
	if time.Now().Unix() > ctx.Exp {
		utils.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTExpired)
		return
	}

	var verified int64 = 1

	//create a user model to update status
	user := db.User{
		Id:     ctx.Id,
		Status: verified,
	}

	//Check if status is 0(unverified) or not
	g, err := user.GetUserById()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}
	if g.Status != 0 {
		utils.ErrorResponse(c, http.StatusForbidden, config.EmailVerifyErr)
		return
	}

	_, err = user.UpdateStatus()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}

	addr := config.INDEX
	utils.RedirectResponse(c, addr)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/middleware"
	"github.com/shiminjia/bookcommunity/model"
	"github.com/shiminjia/bookcommunity/utils"
	"net/http"
)

func Login(c *gin.Context) {
	var u model.Login

	//get user info from request
	if err := c.ShouldBindJSON(&u); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	//todo get user info from db by user name
	//this is a demo now
	var d = model.UserInfoWithId{
		Id:       "1234",
		Email:    "zhangsan@gmail.com",
		Username: "zhangsan",
		Password: "11111111",
	}

	//compare the password with request and db
	if u.Password != d.Password {
		middleware.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError)
		return
	}

	//create a new token
	var ctx = &utils.Context{
		Id:       "d.Id",
		Username: "d.Username",
	}
	t, err := utils.CreateToken(ctx)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, config.CreateTokenError)
		return
	}

	//set the token in data and return it.
	data := &model.Token{Token: t}
	middleware.NormalResponse(c, http.StatusOK, config.OK, data)
}

func Logout(c *gin.Context) {
	middleware.NormalResponse(c, http.StatusOK, config.OK, "")
}

package controller

import (
	"bookcommunity/config"
	"bookcommunity/middleware"
	"bookcommunity/model"
	"bookcommunity/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var u model.CreateUser

	//get user info from request
	if err := c.ShouldBindJSON(&u); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	//create a new token
	var ctx = &utils.Context{
		Id:       "d.Id",
		Email:    "d.Email",
		Username: "d.Username",
	}
	t, err := utils.CreateToken(ctx);
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, config.CreateTokenError)
	}

	//set the token in data and return it.
	data := &model.Token{Token: t}
	middleware.NormalResponse(c, http.StatusOK, config.OK, data)
}

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
		Id: "1234",
		Email: "zhangsan@gmail.com",
		Username: "zhangsan",
		Password: "11111111",
	}

	//compare the password with request and db
	if u.Password != d.Password {
		middleware.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_0)
		return
	}

	//create a new token
	var ctx = &utils.Context{
		Id:       "d.Id",
		Email:    "d.Email",
		Username: "d.Username",
	}
	t, err := utils.CreateToken(ctx);
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, config.CreateTokenError)
		return
	}

	//set the token in data and return it.
	data := &model.Token{Token: t}
	middleware.NormalResponse(c, http.StatusOK, config.OK, data)
}

func Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_1)
		return
	}
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZX"+
		"JuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_2)
		return
	}

	middleware.NormalResponse(c, http.StatusOK, config.OK, "")
}

func GetUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJ"+
		"uYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_2)
		return
	}
	username := c.Param("username")
	data := &model.UserInfo{
		Email:    "testuser1@gmail.com",
		Username: username,
	}
	middleware.NormalResponse(c, http.StatusOK, config.OK, data)
}

func ModifyUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJ"+
		"uYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_2)
		return
	}

	var json model.UserInfo
	if err := c.ShouldBind(&json); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	middleware.NormalResponse(c, http.StatusOK, config.OK, "")
}

func DeleteUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJ"+
		"uYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_2)
		return
	}
	middleware.NormalResponse(c, http.StatusOK, config.OK, "")
}
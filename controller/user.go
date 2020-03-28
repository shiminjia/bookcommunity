package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/middleware"
	"github.com/shiminjia/bookcommunity/model"
	"github.com/shiminjia/bookcommunity/utils"
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
		Username: "d.Username",
	}
	t, err := utils.CreateToken(ctx)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, config.CreateTokenError)
	}

	//set the token in data and return it.
	data := &model.Token{Token: t}
	middleware.NormalResponse(c, http.StatusOK, config.OK, data)
}

func GetUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJ"+
		"uYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_JWTInvalid)
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
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_JWTInvalid)
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
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_JWTInvalid)
		return
	}
	middleware.NormalResponse(c, http.StatusOK, config.OK, "")
}

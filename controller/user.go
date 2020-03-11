package controller

import (
	"bookcommunity/config"
	"bookcommunity/model"
	"bookcommunity/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var json model.CreateUser
	if err := c.ShouldBindJSON(&json); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	data := &model.Token{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1M" +
		"jgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}
	middleware.NormalResponse(c, http.StatusOK, config.OK, data)
}

func Login(c *gin.Context) {
	var json model.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	data := &model.Token{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1M" +
		"jgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}
	middleware.NormalResponse(c, http.StatusOK, config.OK, data)
}

func Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_1)
		return
	}
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZX" +
		"JuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_2)
		return
	}

	middleware.NormalResponse(c, http.StatusOK, config.OK, "")
}

func GetUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJ" +
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
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJ" +
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

func DeleteUser(c *gin.Context){
	token := c.GetHeader("Authorization")
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJ" +
		"uYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_2)
		return
	}
	middleware.NormalResponse(c, http.StatusOK, config.OK, "")
}
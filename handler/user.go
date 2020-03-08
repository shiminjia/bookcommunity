package handler

import (
	"bookcommunity/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var json model.CreateUser
	if err := c.ShouldBindJSON(&json); err != nil {
		resp := &model.ErrorResponse{
			Code:              -1,
			Message:           "400 Bad Request",
			Developer_Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := &model.NormalResponse{
		Code:    0,
		Message: "ok",
		Data: &model.Token{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ",
		},
	}
	c.JSON(http.StatusCreated, resp)
}

func Login(c *gin.Context) {
	var json model.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		resp := &model.ErrorResponse{
			Code:              -1,
			Message:           "400 Bad Request",
			Developer_Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := &model.NormalResponse{
		Code:    0,
		Message: "OK",
		Data: &model.Token{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ",
		},
	}
	c.JSON(http.StatusOK, resp)
}

func Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		resp := &model.ErrorResponse{
			Code:              -1,
			Message:           "unverified error",
			Developer_Message: "JWT is not exist.",
		}
		c.JSON(http.StatusUnauthorized, resp)
		return
	}
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
		resp := &model.ErrorResponse{
			Code:              -1,
			Message:           "unverified error",
			Developer_Message: "JWT is invalid.",
		}
		c.JSON(http.StatusUnauthorized, resp)
		return
	}
	resp := &model.NormalResponse{
		Code:    0,
		Message: "OK",
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}

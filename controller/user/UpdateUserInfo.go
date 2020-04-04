package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/middleware"
	"github.com/shiminjia/bookcommunity/model"
	"net/http"
)

func UpdateUserInfo(c *gin.Context) {
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
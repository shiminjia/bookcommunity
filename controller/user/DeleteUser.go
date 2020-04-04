package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/middleware"
	"net/http"
)

func DeleteUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJ"+
		"uYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_JWTInvalid)
		return
	}
	middleware.NormalResponse(c, http.StatusOK, config.OK, "")
}

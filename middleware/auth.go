package middleware

import (
	"bookcommunity/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context){
		token := c.GetHeader("Authorization")
		if token == "" {
			ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_1)
			c.Abort()
			return
		}
		if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZX"+
			"JuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" {
			ErrorResponse(c, http.StatusBadRequest, config.UnverifiedError_2)
			c.Abort()
			return
		}
	}
}
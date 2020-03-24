package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/utils"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_1)
			c.Abort()
			return
		}

		if bool := utils.VerifyToken(token); bool != true {
			ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_2)
			c.Abort()
			return
		}

		c.Next()
	}
}

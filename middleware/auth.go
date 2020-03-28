package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/utils"
	"net/http"
	"time"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTNonexist)
			c.Abort()
			return
		}

		ctx, err := utils.ParseToken(token)
		if err != nil {
			ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTInvalid)
			c.Abort()
			return
		}

		if ctx.Exp < time.Now().Unix() {
			ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTExpired)
			c.Abort()
			return
		}

		c.Next()
	}
}

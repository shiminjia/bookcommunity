package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/utils"
	"net/http"
	"time"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenBearer := c.GetHeader("Authorization")

		if tokenBearer == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTNonexist)
			c.Abort()
			return
		}

		var token string
		_, err := fmt.Sscanf(tokenBearer, "Bearer %s", &token)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTNoBearer)
			c.Abort()
			return
		}

		ctx, err := utils.ParseToken(token)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTInvalid)
			c.Abort()
			return
		}

		if ctx.Exp < time.Now().Unix() {
			utils.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTExpired)
			c.Abort()
			return
		}

		if ctx.Scope != "login" {
			utils.ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTScopeErr)
			c.Abort()
			return
		}

		//transfer ctx to business login
		c.Set("tokenId", ctx.Id)

		c.Next()
	}
}

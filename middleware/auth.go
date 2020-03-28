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
		tokenbearer := c.GetHeader("Authorization")
		fmt.Println(tokenbearer)
		if tokenbearer == "" {
			ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTNonexist)
			c.Abort()
			return
		}

		var token string
		_, err := fmt.Sscanf(tokenbearer, "Bearer %s", &token)
		fmt.Println(token)
		if err != nil {
			ErrorResponse(c, http.StatusUnauthorized, config.UnverifiedError_JWTNoBearer)
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

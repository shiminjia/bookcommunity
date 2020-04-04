package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/middleware"
	"net/http"
)

func Logout(c *gin.Context) {
	middleware.NormalResponse(c, http.StatusOK, config.OK, "")
}
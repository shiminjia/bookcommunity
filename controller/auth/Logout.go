package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/utils"
	"net/http"
)

func Logout(c *gin.Context) {
	utils.NormalResponse(c, http.StatusOK, config.OK, "")
}
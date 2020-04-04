package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/middleware"
	"github.com/shiminjia/bookcommunity/model"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	userid := c.Param("userid")
	username := "username"

	data := &model.UserInfo{
		Id:       userid,
		Email:    "testuser1@gmail.com",
		Username: username,
	}
	middleware.NormalResponse(c, http.StatusOK, config.OK, data)
}
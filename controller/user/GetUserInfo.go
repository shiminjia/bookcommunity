package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/db"
	"github.com/shiminjia/bookcommunity/model"
	"github.com/shiminjia/bookcommunity/utils"
	"net/http"
	"strconv"
)

func GetUserInfo(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr,10,64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	//create dao object
	dao := db.User{
		Id: uid,
	}

	//get user by id
	u, err := dao.GetUserById()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}
	if u.Id != uid {
		utils.ErrorResponse(c, http.StatusNotFound, config.ErrUserNotFound)
		return
	}
	if u.Status == 0 || u.Status == 8 || u.Status == 9 {
		utils.ErrorResponse(c, http.StatusNotFound, config.ErrUserNotFound)
		return
	}

	data := &model.UserData{
		Id:           u.Id,
		Email:        u.Email,
		Username:     u.Username,
		Icon_url:     u.Icon_url,
		First_name:   u.First_name,
		Middle_name:  u.Middle_name,
		Last_name:    u.Last_name,
		Gender:       u.Gender,
		Birthday:     u.Birthday,
		Location:     u.Location,
		Career:       u.Career,
		Interest:     u.Interest,
		Introduction: u.Introduction,
		Level:        u.Level,
		Status:       u.Status,
		Time:         model.Time{
			CreatedAt:  u.Time.CreatedAt,
			UpdatedAt:  u.Time.UpdatedAt,
			DeletedAt:  u.Time.DeletedAt,
		},
	}
	utils.NormalResponse(c, http.StatusOK, config.OK, data)
}

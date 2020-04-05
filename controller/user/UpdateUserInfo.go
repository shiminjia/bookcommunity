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

func UpdateUserInfo(c *gin.Context) {
	//get uid from path
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	//get id from jwt
	tuid := c.MustGet("tokenId")
	if tuid != uid {
		utils.ErrorResponse(c, http.StatusForbidden, config.Forbidden)
		return
	}

	//bind json to dto
	var dto model.UpdateUser
	if err := c.ShouldBind(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	//check gender is invalid or not
	var genderList = []int64{0,1,2,3,9}
	var IsInList = false
	for _,k := range genderList {
		if k == dto.Gender {
			IsInList = true
			break;
		}
	}
	if IsInList != true {
		utils.ErrorResponse(c, http.StatusBadRequest, config.ErrGender)
		return
	}

	//todo check dto.Icon_url is invalid
	//if you upload this icon at update-user-info page, the scope of this icon is user_icon
	//check the scope dto.Icon_url of is user_icon or not

	//create dao
	var dao = db.User{
		Id:           uid,
		Username:     dto.Username,
		Icon_url:     dto.Icon_url,
		First_name:   dto.First_name,
		Middle_name:  dto.Middle_name,
		Last_name:    dto.Last_name,
		Gender:       dto.Gender,
		Birthday:     dto.Birthday,
		Location:     dto.Location,
		Career:       dto.Career,
		Interest:     dto.Interest,
		Introduction: dto.Introduction,
	}

	//update user by id
	_, err = dao.UpdateUserById()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}

	utils.NormalResponse(c, http.StatusOK, config.OK, "")
}
package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/db"
	"github.com/shiminjia/bookcommunity/model"
	"github.com/shiminjia/bookcommunity/utils"
	"net/http"
)

func Login(c *gin.Context) {
	var dto model.Login

	//get email and password from request
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	//todo get user info from db by user name
	dkStr, err := utils.Encrypt(dto.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.InternalServerError)
		return
	}

	var dao = db.User{
		Email:    dto.Email,
		Password: dkStr,
	}

	v, err := dao.VerifyByEmailAndPassWord()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}

	if len(*v) != 1 {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.LOGINFailure)
		return
	}

	//create a new token
	var ctx = &utils.Context{
		Id:       (*v)[0].Id,
		Username: (*v)[0].Username,
		Scope:    "login",
	}

	var eft = config.EXPIRATION_TIME_LOGIN

	t, err := utils.CreateToken(ctx, eft)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.CreateTokenError)
		return
	}

	//set the token in data and return it.
	data := &model.LoginData{
		UserId: (*v)[0].Id,
		Token:  t,
	}
	utils.NormalResponse(c, http.StatusOK, config.OK, data)
}

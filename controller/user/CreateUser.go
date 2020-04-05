package user

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/db"
	"github.com/shiminjia/bookcommunity/model"
	"github.com/shiminjia/bookcommunity/utils"
	"golang.org/x/crypto/scrypt"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var dto model.CreateUser

	//get user info from request
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	//Validate Struct
	_, err := govalidator.ValidateStruct(&dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, config.ValidError)
		return
	}

	//Encrypt password and change byte[] to string
	dk, err := scrypt.Key([]byte(dto.Password), []byte(config.SALT), 16384, 8, 1, 32)
	dkStr := fmt.Sprintf("%x", dk)

	//create a new user model
	dao := db.User{
		Email:    dto.Email,
		Password: dkStr,
		Username: dto.Username,
	}

	//ensure email and username unique
	res, err := dao.Verify()
	if err != nil && err.Error() != "record not found" {
		fmt.Println(err.Error())
		utils.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}
	if res.Email == dao.Email {
		utils.ErrorResponse(c, http.StatusBadRequest, config.EmailUniqueErr)
		return
	}
	if res.Username == dao.Username {
		utils.ErrorResponse(c, http.StatusBadRequest, config.UsernameUniqueErr)
		return
	}

	//insert user model into db
	res, err = dao.Create()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}

	//create a new token
	var ctx = &utils.Context{
		Id:       res.Id,
		Username: "d.Username",
		Scope:    "email",
	}
	var eft = config.EXPIRATION_TIME_EMAIL
	t, err := utils.CreateToken(ctx, eft)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, config.CreateTokenError)
		return
	}

	//send email

	if config.MODE == "dev" || config.MODE == "test" || config.MODE == "" {
		fmt.Println("http://localhost:8080/auth?token=" + t)
	} else {
		utils.SendEmail()
	}

	//response
	data := &model.UserId{UserId: ctx.Id}
	utils.NormalResponse(c, http.StatusOK, config.OK, data)

}

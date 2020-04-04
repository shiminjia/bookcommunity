package user

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/db"
	"github.com/shiminjia/bookcommunity/middleware"
	"github.com/shiminjia/bookcommunity/model"
	"github.com/shiminjia/bookcommunity/utils"
	"golang.org/x/crypto/scrypt"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var dto model.CreateUser

	//get user info from request
	if err := c.ShouldBindJSON(&dto); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.ErrBind)
		return
	}

	//Validate Struct
	_, err := govalidator.ValidateStruct(&dto)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.ValidError)
		return
	}

	//Encrypt password and change byte[] to string
	dk, err := scrypt.Key([]byte(dto.Password), []byte(config.SALT), 16384, 8, 1, 32)
	dkstr := fmt.Sprintf("%x",dk)

	//create a new user model
	dao := db.User  {
		Email:      dto.Email,
		Password:   dkstr,
		Username:   dto.Username,
	}

	//ensure email and username unique
	v ,err := dao.Verify()
	if err != nil && err.Error() != "record not found"  {
		fmt.Println(err.Error())
		middleware.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}
	if v.Email == dao.Email {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.EmailUniqueErr)
		return
	}
	if v.Username == dao.Username {
		middleware.ErrorResponse(c, http.StatusBadRequest, config.UsernameUniqueErr)
		return
	}

	//insert user model into db
	v, err = dao.Create()
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, config.DBErr)
		return
	}

	//create a new token
	IdUint := strconv.Itoa(int(v.Id.Id))
	var ctx = &utils.Context{
		Id:       IdUint,
		Username: "d.Username",
		Scope:    "email",
	}
	t, err := utils.CreateToken(ctx)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, config.CreateTokenError)
		return
	}
	fmt.Println("http://localhost:8080/auth?token=" + t)

	//send email
	//utils.SendEmail()

	//response
	data := &model.UserId{UserId: IdUint}
	middleware.NormalResponse(c, http.StatusOK, config.OK, data)

}





package model

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)

	const Username = "^[a-zA-Z0-9]+$"

	var rxUsername = regexp.MustCompile(Username)

	govalidator.TagMap["username"] = govalidator.Validator(func(str string) bool {
		return rxUsername.MatchString(str)
	})
}

type CreateUser struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"            valid:"length(8|100), email, required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"   valid:"length(8|30), username, required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"   valid:"length(8|30), alphanum, required"`
}

type Login struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"            valid:"length(8|100), email, required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"   valid:"length(8|100), email, required"`
}

type UserInfo struct {
	Id       string `form:"id" json:"id" xml:"id"  binding:"required"                     valid:"length(8|100), email, required"`
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"            valid:"length(8|100), email, required"     `
	Username string `form:"username" json:"username" xml:"username"  binding:"required"   valid:"length(8|100), email, required"`
}

type UserInfoWithId struct {
	Id       string `form:"id" json:"id" xml:"id"  binding:"required"                     valid:"length(8|100), email, required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"   valid:"length(8|100), email, required"`
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"            valid:"length(8|100), email, required"       `
	Password string `form:"password" json:"password" xml:"password"  binding:"required"   valid:"length(8|100), email, required"`
}


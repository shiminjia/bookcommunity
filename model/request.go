package model

import "time"

type CreateUser struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"            valid:"length(8|100), email, required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"   valid:"length(8|30), username, required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"   valid:"length(8|30), alphanum, required"`
}

type Login struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"            valid:"length(8|100), email, required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"   valid:"length(8|100), email, required"`
}

//type UserInfo struct {
//	Id       int64 `form:"id" json:"id" xml:"id"  binding:"required"                     valid:"length(8|100), email, required"`
//	Email    string `form:"email" json:"email" xml:"email"  binding:"required"            valid:"length(8|100), email, required"     `
//	Username string `form:"username" json:"username" xml:"username"  binding:"required"   valid:"length(8|100), email, required"`
//}

//type UserInfoWithId struct {
//	Id       int64 `form:"id" json:"id" xml:"id"  binding:"required"                     valid:"length(8|100), email, required"`
//	Username string `form:"username" json:"username" xml:"username"  binding:"required"   valid:"length(8|100), email, required"`
//	Email    string `form:"email" json:"email" xml:"email"  binding:"required"            valid:"length(8|100), email, required"       `
//	Password string `form:"password" json:"password" xml:"password"  binding:"required"   valid:"length(8|100), email, required"`
//}

type UpdateUser struct {
	Username     string `form:"username" json:"username" xml:"username"          valid:"length(8|30), username, optional"`
	Icon_url     string `form:"icon_url" json:"icon_url" xml:"icon_url"          valid:"length(0|100), url, optional"`
	First_name   string `form:"first_name" json:"first_name" xml:"first_name"    valid:"length(0|30), username, optional"`
	Middle_name  string `form:"middle_name" json:"middle_name" xml:"middle_name" valid:"length(0|30), username, optional"`
	Last_name    string `form:"last_name" json:"last_name" xml:"last_name"       valid:"length(0|30), username, optional"`
	Gender       int64  `form:"gender" json:"gender" xml:"gender"                valid:"range(0|9), utfnumeric, optional"`
	Birthday     *time.Time `form:"birthday" json:"birthday" xml:"birthday"          valid:"length(25|25), utfletternum, optional"`
	Location     string `form:"location" json:"location" xml:"location"          valid:"length(0|50), utfletternum, optional"`
	Career       string `form:"career" json:"career" xml:"career"                valid:"length(0|100), utfletternum, optional"`
	Interest     string `form:"interest" json:"interest" xml:"interest"          valid:"length(0|100), utfletternum, optional"`
	Introduction string `form:"introduction" json:"introduction" xml:"introduction" valid:"length(0|1000), utfletternum, optional"`
}

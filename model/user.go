package model

type CreateUser struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

type Login struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

type UserInfo struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
}

type UserInfoWithId struct {
	Id       string `form:"id" json:"id" xml:"id"  binding:"required"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

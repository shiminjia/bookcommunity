package model

import "time"

//NormalResponse
type NormalResponse struct {
	Code    int         `form:"code" json:"code" xml:"code"`
	Message string      `form:"message" json:"message" xml:"message"`
	Data    interface{} `form:"data" json:"data" xml:"data"`
}

type UserId struct {
	UserId int64 `form:"user_id" json:"user_id" xml:"user_id"`
}

type Token struct {
	Token string `form:"token" json:"token" xml:"token"`
}

type LoginData struct {
	UserId int64 `form:"user_id" json:"user_id" xml:"user_id"`
	Token string `form:"token" json:"token" xml:"token"`
}

type UserData struct {
	Id           int64      `form:"id" json:"id" xml:"id"`
	Email        string    `form:"email" json:"email" xml:"email"`
	Username     string    `form:"username" json:"username" xml:"username"`
	Icon_url     string    `form:"icon_url" json:"icon_url" xml:"icon_url"`
	First_name   string    `form:"first_name" json:"first_name" xml:"first_name"`
	Middle_name  string    `form:"middle_name" json:"middle_name" xml:"middle_name"`
	Last_name    string    `form:"last_name" json:"last_name" xml:"last_name"`
	Gender       int64      `form:"gender" json:"gender" xml:"gender"`
	Birthday     *time.Time  `form:"birthday" json:"birthday" xml:"birthday"`
	Location     string     `form:"location" json:"location" xml:"location"`
	Career       string     `form:"career" json:"career" xml:"career"`
	Interest     string     `form:"interest" json:"interest" xml:"interest"`
	Introduction string     `form:"introduction" json:"introduction" xml:"introduction"`
	Level        int64      `form:"level" json:"level" xml:"level"`
	Status       int64      `form:"status" json:"status" xml:"status"`
	Time
}

type Time struct {
	CreatedAt time.Time    `form:"create_at" json:"create_at" xml:"create_at"`
	UpdatedAt time.Time    `form:"update_at" json:"update_at" xml:"update_at"`
	DeletedAt *time.Time   `form:"delete_at" json:"delete_at" xml:"delete_at"`
}

//ErrorResponse
type ErrorResponse struct {
	Code              int         `form:"code" json:"code" xml:"code"`
	Message           string      `form:"message" json:"message" xml:"message"`
	Developer_Message interface{} `form:"develop_message" json:"develop_message" xml:"develop_message"`
}

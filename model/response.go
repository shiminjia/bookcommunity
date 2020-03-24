package model

//NormalResponse
type NormalResponse struct {
	Code    int         `form:"code" json:"code" xml:"code"`
	Message string      `form:"message" json:"message" xml:"message"`
	Data    interface{} `form:"data" json:"data" xml:"data"`
}

type Token struct {
	Token string `form:"token" json:"token" xml:"token"`
}

//ErrorResponse
type ErrorResponse struct {
	Code              int         `form:"code" json:"code" xml:"code"`
	Message           string      `form:"message" json:"message" xml:"message"`
	Developer_Message interface{} `form:"develop_message" json:"develop_message" xml:"develop_message"`
}

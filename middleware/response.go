package middleware

import (
	"bookcommunity/config"
	"bookcommunity/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NormalResponse(c *gin.Context, httpstatus int, code int, data interface{}) {
	resp := &model.NormalResponse{
		Code:    code,
		Message: http.StatusText(httpstatus),
		Data:    data,
	}
	c.JSON(httpstatus, resp)
}

func ErrorResponse(c *gin.Context, httpstatus int, code int){
	resp := &model.ErrorResponse{
		Code:              code,
		Message:           http.StatusText(httpstatus),
		Developer_Message: config.CodeText(code),
	}
	c.JSON(httpstatus, resp)
}
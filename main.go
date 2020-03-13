package main

import (
	"bookcommunity/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users", controller.CreateUser)
	router.PUT("/users", controller.Login)
	router.DELETE("/users", controller.Logout)
	router.GET("/users/:username", controller.GetUserInfo)
	router.POST("/users/:username", controller.ModifyUserInfo)
	router.DELETE("/users/:username", controller.DeleteUser)

	router.Run(":8080")
}


//todo 日志答应
//拦截请求

//日志，链路
//

//数据库的调用 mysql / redis
//

//jwt 与中间件
//
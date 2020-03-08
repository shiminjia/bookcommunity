package main

import (
	"bookcommunity/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users", handler.CreateUser)
	router.PUT("/users/login", handler.Login)
	router.GET("/users/logout", handler.Logout)
	//router.GET("/users/:username", GetUserInfo)
	//router.POST("/users/:username", ModifyUserInfo)
	//router.DELETE("/users/:username", DeleteUser)

	router.Run(":8080")
}

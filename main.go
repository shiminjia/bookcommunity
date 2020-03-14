package main

import (
	"bookcommunity/controller"
	"bookcommunity/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	//use middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//the router that no need to auth
	router.POST("/users", controller.CreateUser)
	router.PUT("/users", controller.Login)

	//declaim group "/"
	authorized := router.Group("/")

	//group "/" need to auth
	authorized.Use(middleware.AuthRequired())
	{
		authorized.DELETE("/users", controller.Logout)
		authorized.GET("/users/:username", controller.GetUserInfo)
		authorized.POST("/users/:username", controller.ModifyUserInfo)
		authorized.DELETE("/users/:username", controller.DeleteUser)
	}

	router.Run(":8080")
}

//
// 1.output
// output the log at the console only when mode is development.
// output the log at the console and log file when mode is test/debug.
// output the log at log file when mode is production.
//
// 2.config file
// a config need to include three patten of configuraion which is development, test/debug and production.
//
// 3.jwt middleware
//
// 4.use database
//
// 5.Intercept invalid access
//


package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/controller"
	"github.com/shiminjia/bookcommunity/middleware"
	"io"
	"os"
)

func getConfig() (string, string) {
	mode := config.MODE
	if mode == "" {
		mode = "dev"
	}

	addr := config.ADDR
	if addr == "" {
		addr = ":8080"
	}

	return mode, addr
}

func main() {

	router := gin.New()

	mode, addr := getConfig()

	//use middleware
	if mode == "dev" {
		router.Use(gin.Logger())
	} else if mode == "test" {
		router.Use(gin.Logger())
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	} else if mode == "prod" {
		gin.DisableConsoleColor()
		//todo need to change the name of log like yyyyMMdd.log
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	} else {
		router.Use(gin.Logger())
	}

	router.Use(gin.Recovery())

	//declaim the router that no need to auth
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

	router.Run(addr)
}

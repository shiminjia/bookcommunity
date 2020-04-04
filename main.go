package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/bookcommunity/config"
	"github.com/shiminjia/bookcommunity/controller/auth"
	"github.com/shiminjia/bookcommunity/controller/user"
	"github.com/shiminjia/bookcommunity/db"
	"github.com/shiminjia/bookcommunity/middleware"
	"io"
	"os"
)

func main() {

	router := gin.New()

	//get addr from config
	_, addr := getConfig()

	//DB access
	db.Init()

	//use middleware
	middlewareRegister(router)

	//router registration
	//declaim the router that no need to auth
	router.GET("/auth", auth.EmailVerification)
	router.POST("/auth", auth.Login)

	router.POST("/users", user.CreateUser)
	router.GET("/users/:userid", user.GetUserInfo)

	//declaim group "/"
	authorized := router.Group("/")

	//group "/" need to auth
	authorized.Use(middleware.AuthRequired())
	{
		authorized.DELETE("/auth", auth.Logout)

		authorized.PUT("/users/:userid", user.UpdateUserInfo)
		//authorized.GET("/users/:username/password", controller.SendChangePasswordMail)
		//authorized.PUT("/users/:username/password", controller.UpdatePassword)
		//authorized.DELETE("/users/:username", controller.DeleteUser)
	}

	router.Run(addr)
}

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

func middlewareRegister(router *gin.Engine) {

	//get value from config
	mode, _ := getConfig()

	//use Logger()
	switch mode {
	case "dev":
		router.Use(gin.Logger())
	case "test":
		router.Use(gin.Logger())
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	case "prod":
		gin.DisableConsoleColor()
		//todo need to change the name of log like yyyyMMdd.log
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	default:
		router.Use(gin.Logger())
	}

	//use Recovery()
	router.Use(gin.Recovery())
}

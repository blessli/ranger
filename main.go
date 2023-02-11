package main

import (
	handler "github.com/blessli/ranger/router"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.New()

func main() {
	router.POST("/dev-api/api/sys/user/login", handler.HandleLogin)
	router.GET("/dev-api/api/sys/user/currentUser", handler.HandleUserInfo)
	router.POST("/dev-api/api/sys/user/list", handler.HandleUserList)
	router.POST("/dev-api/api/sys/role/list", handler.HandleRoleList)
	err := router.Run(":8888")
	if err != nil {
		panic(err)
	}
}

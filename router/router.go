package router

import (
	"github.com/blessli/ranger/services"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

var loginService = services.NewLoginService()

func HandleLogin(c *gin.Context) {
	logx.Info("HandleLogin start")
	rsp, err := loginService.Login("admin", "123456")
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, rsp)
}

func HandleUserInfo(c *gin.Context) {
	logx.Info("HandleUserInfo start")
	rsp, err := loginService.UserInfo(1)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, rsp)
}

func HandleUserList(c *gin.Context) {
	logx.Info("HandleUserList start")
	req := &services.ListUserReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	rsp, err := loginService.UserList(req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, rsp)
}

func HandleRoleList(c *gin.Context) {
	logx.Info("HandleRoleList start")
	req := &services.ListRoleReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	rsp, err := loginService.RoleList(req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, rsp)
}
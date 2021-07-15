package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
)

func RouteUser(rg *gin.RouterGroup) {
	Router := rg.Group("user")
	{
		Router.POST("login", v1.Login)
		//Router.POST("out", v1.UserLogout)
	}
}
package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
)

func RouteUser(rg *gin.RouterGroup) {
	router := rg.Group("user")
	{
		router.POST("login", v1.Login)
		router.POST("logout", v1.Logout)
		router.POST("register", v1.Register)
		router.POST("captcha", v1.Captcha)
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
	"github.com/saicem/todo/middleware"
)

func RouteTodoGroup(rg *gin.RouterGroup) {
	Router := rg.Group("todo_group", middleware.Authentication)
	{
		Router.POST("add", v1.TodoGroupAdd)
		Router.GET("list", v1.TodoGroupGet)
		Router.PUT(":id", v1.TodoGroupPut)
		Router.DELETE(":id", v1.TodoGroupDelete)
	}
}

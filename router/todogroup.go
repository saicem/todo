package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
	"github.com/saicem/todo/middleware"
)

func RouteTodoGroup(rg *gin.RouterGroup) {
	router := rg.Group("todo_group", middleware.Authentication)
	{
		router.POST("add", v1.TodoGroupAdd)
		router.GET("list", v1.TodoGroupGet)
		router.PUT(":id", v1.TodoGroupPut)
		router.DELETE(":id", v1.TodoGroupDelete)
	}
}

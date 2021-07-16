package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
	"github.com/saicem/todo/middleware"
)

func RouteTodo(rg *gin.RouterGroup) {
	TodoRouter := rg.Group("todo", middleware.Authentication)
	{
		TodoRouter.POST("add", v1.TodoAdd)
		TodoRouter.GET("list", v1.TodoGet)
		TodoRouter.PUT(":id", v1.TodoPut)
		TodoRouter.DELETE(":id", v1.TodoDelete)
	}
}

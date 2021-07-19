package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
	"github.com/saicem/todo/middleware"
)

func RouteTodo(rg *gin.RouterGroup) {
	router := rg.Group("todo", middleware.Authentication)
	{
		router.POST("add", v1.TodoAdd)
		router.GET("list", v1.TodoGet)
		router.PUT(":id", v1.TodoPut)
		router.DELETE(":id", v1.TodoDelete)
	}
}

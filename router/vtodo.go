package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
	"github.com/saicem/todo/middleware"
)

func RouteTodo(rg *gin.RouterGroup) {
	TodoRouter := rg.Group("vtodo", middleware.Authentication)
	{
		TodoRouter.POST("add", v1.TodoAdd)
		TodoRouter.GET("get", v1.TodoGet)
		TodoRouter.POST("complete", v1.TodoComplete)
		TodoRouter.POST("update", v1.TodoUpdate)
		TodoRouter.DELETE("delete", v1.TodoDelete)
	}
}

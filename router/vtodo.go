package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
)

func RouteTodo(rg *gin.RouterGroup) {
	TodoRouter := rg.Group("vtodo")
	{
		TodoRouter.POST("add", v1.TodoAdd)
		TodoRouter.POST("complete", v1.TodoComplete)
		TodoRouter.POST("change", v1.TodoChange)
		TodoRouter.DELETE("delete", v1.TodoDelete)
	}
}

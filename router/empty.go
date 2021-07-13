package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
)

// RouteEmpty
// @Summary ping
// @Description ping
func RouteEmpty(rg *gin.RouterGroup) {
	EmptyRouter := rg.Group("")
	{
		EmptyRouter.GET("ping", v1.Ping)
	}
}

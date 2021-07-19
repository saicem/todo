package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/todo/api/v1"
)

// RouteEmpty
// @Summary ping
// @Description ping
func RouteEmpty(rg *gin.RouterGroup) {
	router := rg.Group("")
	{
		router.GET("ping", v1.Ping)
	}
}

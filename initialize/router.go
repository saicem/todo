package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/config"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	engine := gin.New()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("open swagger UI http://localhost:%s/swagger/index.html\n", config.ProjectPort)
	BasicGroup := engine.Group("api/v1", middleware.Cors)
	router.RouteEmpty(BasicGroup)
	router.RouteTodo(BasicGroup)
	router.RouteUser(BasicGroup)

	return engine
}

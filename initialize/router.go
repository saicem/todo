package initialize

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/config"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	initSwagger(engine)
	basicGroup := engine.Group("api/v1", middleware.Cors)
	router.RouteEmpty(basicGroup)
	router.RouteTodo(basicGroup)
	router.RouteUser(basicGroup)
	router.RouteTodoGroup(basicGroup)

	return engine
}

func initSwagger(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Printf("open swagger UI http://localhost:%s/swagger/index.html\n", config.ProjectPort)
}

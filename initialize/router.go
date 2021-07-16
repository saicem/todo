package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/config"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"os/exec"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	initSwagger(engine)
	BasicGroup := engine.Group("api/v1", middleware.Cors)
	router.RouteEmpty(BasicGroup)
	router.RouteTodo(BasicGroup)
	router.RouteUser(BasicGroup)
	router.RouteTodoGroup(BasicGroup)

	return engine
}

func initSwagger(engine *gin.Engine) {
	cmd := exec.Command("swag", "init")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		// todo 更高的错误等级
		panic(err)
	}
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("open swagger UI http://localhost:%s/swagger/index.html\n", config.ProjectPort)
}

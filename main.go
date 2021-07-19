package main

import (
	"log"

	"github.com/saicem/todo/config"
	_ "github.com/saicem/todo/docs"
	"github.com/saicem/todo/initialize"
)

// @title vTodo Api
// @version 0.1
// @description This is a vTodo Api

// @contact.name saicem
// @contact.url  www.saicem.top
// @contact.email saicem@foxmail.com

// @host localhost:9101
// @BasePath /api/v1

func main() {
	initialize.InitMysql()
	initialize.InitRedis()
	engine := initialize.InitRouter()
	err := engine.Run(":" + config.ProjectPort)
	if err != nil {
		log.Panicln(err)
	}
}

// redis todo 该怎么用？

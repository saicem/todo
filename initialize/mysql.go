package initialize

import (
	"log"
	"os"

	"github.com/saicem/todo/global"
	"github.com/saicem/todo/model/todomodel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	log.Println("init mysql ... ... ")
	NewConn()
	migrate()
}

func migrate() {
	err := global.Mysql.AutoMigrate(
		&todomodel.User{},
		&todomodel.TodoGroup{},
		&todomodel.TodoItem{},
	)
	if err != nil {
		log.Println(err)
		log.Println("迁移数据库失败")
		os.Exit(0)
	}
}

func NewConn() {
	config := global.Config
	db, err := gorm.Open(mysql.Open(config.Mysql.Dsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		global.Mysql = db
		log.Println("new mysql conn ...")
	}
}

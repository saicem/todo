package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/saicem/todo/config"
	"gorm.io/gorm"
)

var (
	Mysql  *gorm.DB
	Redis  *redis.Pool
	Config *config.Config
)

func init() {
	Config = config.Get()
}

package global

import (
	"github.com/saicem/todo/config"
	"gorm.io/gorm"
)

var (
	Mysql  *gorm.DB
	Config *config.Config
)

func init() {
	Config = config.Get()
}

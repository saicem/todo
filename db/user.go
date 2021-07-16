package db

import (
	"github.com/saicem/todo/global"
	"github.com/saicem/todo/model/todomodel"
)

func SearchUser(userName string, password string) (*todomodel.User, bool) {
	var user todomodel.User
	res := global.Mysql.Where("user_id = ? AND password = ?", userName, password).First(&user)
	if res.Error != nil {
		return nil, false
	}
	return &user, true
}

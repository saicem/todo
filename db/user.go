package db

import (
	"regexp"

	"github.com/saicem/todo/global"
	"github.com/saicem/todo/model/todomodel"
)

func SearchUser(userName string, password string) (*todomodel.User, bool) {
	isMatch, _ := regexp.MatchString(".+@.+", userName)
	var user *todomodel.User
	if isMatch {
		res := global.Mysql.Where("email = ? AND password = ?", userName, password).First(&user)
		if res.Error != nil {
			return nil, false
		}
	} else {
		err := global.Mysql.Where("user_id = ? AND password = ?", userName, password).First(&user)
		if err != nil {
			return nil, false
		}
	}
	return user, true
}

func NewUser(email string, password string) *todomodel.User {
	var user todomodel.User
	res := global.Mysql.Where("email = ?", email).First(&user)
	if res.Error == nil {
		// 邮箱已注册
		return nil
	}
	user = todomodel.User{
		Email:    email,
		Password: password,
	}
	global.Mysql.Create(&user)
	return &user
}

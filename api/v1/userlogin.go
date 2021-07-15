package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/saicem/todo/global"
	"github.com/saicem/todo/model/response"
	"github.com/saicem/todo/model/todomodel"
	"math/rand"
	"net/http"
	"time"
)

// Login
// @Summary Login
// @Description 用户登录
// @Param uid query string true "用户ID"
// @Param password query string true "密码"
// @Router /user/login [post]
// @Success 200 object response.Response
func Login(c *gin.Context) {
	// 获取参数
	userName := c.Query("uid")
	password := c.Query("password")
	if isValid := SearchUser(userName, password); !isValid {
		c.JSON(http.StatusOK, response.Response{Status: response.ERROR, Message: "未通过验证"})
	} else {
		sessionId := RandString(50)
		maxAge := global.Config.Session.MaxAge
		domain := global.Config.Session.Domain
		c.SetCookie("SESSIONID", sessionId, maxAge, "/", domain, false, true)
		r := global.Redis.Get()
		if _, err := r.Do("SET", sessionId, 1, "EX", maxAge); err != nil {
			panic("发送")
		}
		defer func(r redis.Conn) {
			err := r.Close()
			if err != nil {
				panic("关不掉？？")
			}
		}(r)
		c.JSON(http.StatusOK, response.Response{Status: response.OK, Message: "登录成功"})
	}
}

func SearchUser(userName string, password string) bool {
	db := global.Mysql
	var user todomodel.User
	// todo 判断存在的更好的方式
	res := db.Where("uid = ? AND password = ?", userName, password).First(&user)
	if res.Error != nil {
		return false
	}
	return true
}

func RandString(length int) string {
	str := "0123456789QWERTYUIOPASDFGHJKLZXCVBNM"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

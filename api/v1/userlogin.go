package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/global"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
	"math/rand"
	"net/http"
	"time"
)

// Login
// @Summary Login
// @Description 用户登录
// @Param json body request.LoginForm true "登录表单"
// @Router /user/login [post]
// @Success 200 object response.Response
func Login(c *gin.Context) {
	// 获取参数
	var loginForm request.LoginForm
	if err := c.BindJSON(&loginForm); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user, isValid := db.SearchUser(loginForm.Username, loginForm.Password)
	if !isValid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	sessionId := randString(50)
	maxAge := global.Config.Session.MaxAge
	domain := global.Config.Session.Domain
	c.SetCookie("SESSIONID", sessionId, maxAge, "/", domain, false, true)
	db.SetSession(sessionId, user.UserId)
	c.JSON(http.StatusOK, response.Response{Msg: "登录成功"})
}

func randString(length int) string {
	str := "0123456789QWERTYUIOPASDFGHJKLZXCVBNM"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

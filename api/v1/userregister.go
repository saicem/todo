package v1

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
)

// Register
// @Summary Register
// @Description 注册
// @Param json body request.RegisterForm true "注册表单"
// @Router /user/register [post]
// @Success 200 object response.Response
func Register(c *gin.Context) {
	var registerForm request.RegisterForm
	if err := c.BindJSON(&registerForm); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	isOk := db.VerifyCaptcha(registerForm.Captcha, registerForm.Email)
	if !isOk {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Msg:  "success",
		Data: db.NewUser(registerForm.Email, registerForm.Password),
	})
}

// Captcha
// @Summary Captcha
// @Description 获取验证码
// @Param email query string true "邮箱地址"
// @Router /user/captcha [post]
// @Success 200 object response.Response
func Captcha(c *gin.Context) {
	// 获取IP
	// clientIp := c.ClientIP()
	toEmailAddress := c.Query("email")
	captcha := randNum(6)
	if err := middleware.SendMail(toEmailAddress, "TODO 验证码", captcha); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	db.SetCaptcha(captcha, toEmailAddress)
	c.JSON(http.StatusOK, response.Response{Msg: "success"})
}

func randNum(length int) string {
	bytes := []byte("0123456789")
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

package middleware

import (
	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	// todo 记得待会设置一下
	c.Header("Access-Control-Allow-Origin", "*")
	return
}

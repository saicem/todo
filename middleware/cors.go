package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/global"
)

func Cors(c *gin.Context) {
	cors := global.Config.Web.Cors
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", cors)
}

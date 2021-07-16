package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/saicem/todo/global"
	"net/http"
)

func Authentication(c *gin.Context) {
	if sessionId, err := c.Cookie("SESSIONID"); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		userId, isValid := searchSession(sessionId)
		if !isValid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("userId", userId)
	}
	return
}

func searchSession(sessionId string) (int, bool) {
	// todo 不能整个redis全给存这个 需要优化存储策略
	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic("关不掉？？")
		}
	}(r)
	if userId, err := redis.Int(r.Do("GET", sessionId)); err == nil {
		return userId, true
	}
	return 0, false
}

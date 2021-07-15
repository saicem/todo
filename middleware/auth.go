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
		if isValid := SearchSession(sessionId); !isValid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
	return
}

func SearchSession(sessionId string) bool {
	// todo 不能整个redis全给存这个 需要优化存储策略

	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic("关不掉？？")
		}
	}(r)
	if _, err := r.Do("GET", sessionId); err == nil {
		return true
	}
	return false
}

func GetUserIdFromCookie(sessionId string) (uint, bool) {
	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic("关不掉？？")
		}
	}(r)
	// 不能直接转换为 uint 所以先 int 再 uint
	if userId, err := redis.Int(r.Do("GET", sessionId)); err == nil {
		return uint(userId), true
	}
	return 0, false
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
)

func Authentication(c *gin.Context) {
	if sessionId, err := c.Cookie("SESSIONID"); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		userId, isValid := db.SearchSession(sessionId)
		if !isValid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("userId", userId)
	}
}

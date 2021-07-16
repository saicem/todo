package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/model/request"
	"net/http"
)

func Register(c *gin.Context) {
	var loginForm request.LoginForm
	if err := c.BindJSON(&loginForm); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

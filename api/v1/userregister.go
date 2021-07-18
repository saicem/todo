package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/model/request"
)

func Register(c *gin.Context) {
	var loginForm request.LoginForm
	if err := c.BindJSON(&loginForm); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

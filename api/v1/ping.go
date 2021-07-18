package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping
// @Summary ping
// @Description 连接测试
// @Router /ping [get]
// @Success 200 string pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "PONG")
}

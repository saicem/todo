package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

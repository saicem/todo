package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

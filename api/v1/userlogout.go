package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/model/response"
)

// Logout
// @Summary Logout
// @Description 用户登出
// @Router /user/logout [post]
// @Success 200 object response.Response
func Logout(c *gin.Context) {
	sessionId, _ := c.Cookie("SESSIONID")
	db.DELSession(sessionId)
	c.JSON(http.StatusOK, response.Response{
		Msg: "Success",
	})
}

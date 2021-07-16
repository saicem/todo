package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoGroupGet
// @Summary TodoGroupGet
// @Description 查看已有的TodoGroup
// @Router /todo_group/list [get]
// @Success 200 object response.Response
func TodoGroupGet(c *gin.Context) {
	cookie, _ := c.Cookie("SESSIONID")
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	todoGroups := db.GetTodoGroups(userId)
	c.JSON(http.StatusOK, response.Response{
		Msg:  "success",
		Data: todoGroups,
	})
}

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoGet
// @Summary TodoGet
// @Description 查看已有的todo
// @Router /todo/list [get]
// @Success 200 object response.Response
func TodoGet(c *gin.Context) {
	// 从cookie 获取 userId
	cookie, _ := c.Cookie("SESSIONID")
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	todoItems := db.GetTodoItems(userId)
	c.JSON(http.StatusOK, response.Response{
		Msg:  "success",
		Data: todoItems,
	})
	return
}

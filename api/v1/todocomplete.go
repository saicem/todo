package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoComplete
// @Summary TodoComplete
// @Description 一个todo完成了
// @Router /vtodo/complete [post]
// @Param todo_id query int true "todo的ID"
// @Success 200 object response.Response
func TodoComplete(c *gin.Context) {
	todoId, hasId := c.GetQuery("todo_id")
	if hasId == false {
		c.JSON(http.StatusOK, response.Response{
			Status:  response.ERROR,
			Message: "缺少参数",
			Data:    nil,
		})
		return
	}
	// 从cookie 获取 uid
	cookie, _ := c.Cookie("SESSIONID")
	uid, _ := middleware.GetUserIdFromCookie(cookie)
	// 查询数据库
	isSuccess := db.CompleteTodoItem(uid, todoId)
	if !isSuccess {
		c.JSON(http.StatusOK, response.Response{
			Status:  response.ERROR,
			Message: "错误的 todo_id",
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Status:  response.OK,
		Message: "ok",
		Data:    nil,
	})
}

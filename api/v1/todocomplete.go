package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/global"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/response"
	"github.com/saicem/todo/model/todomodel"
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
	// 从cookie 获取 userId
	cookie, _ := c.Cookie("SESSIONID")
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	// 查询数据库
	db := global.Mysql
	var todoItem todomodel.TodoItem
	db.Find(&todoItem, "uid = ? AND id = ?", userId, todoId)
	if &todoItem == nil {
		c.JSON(http.StatusOK, response.Response{
			Status:  response.ERROR,
			Message: "错误的 todo_id",
		})
	}
	// 修改数据库
	todoItem.Complete = true
	db.Save(&todoItem)
	// 返回
	c.JSON(http.StatusOK, response.Response{
		Status:  response.OK,
		Message: "ok",
		Data:    nil,
	})
}

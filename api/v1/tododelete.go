package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/global"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/response"
	"github.com/saicem/todo/model/todomodel"
	"net/http"
)

// TodoDelete
// @Summary TodoDelete
// @Description 删除一个todo
// @Router /vtodo/delete [delete]
// @Param todo_id query int true "todo的id"
// @Success 200 object response.Response
func TodoDelete(c *gin.Context) {
	db := global.Mysql
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
	var todoItem todomodel.TodoItem
	db.Where("uid = ? AND id = ?", userId, todoId).Find(todoItem)
	db.Delete(&todoItem)
	c.JSON(http.StatusOK, response.Response{
		Status:  response.OK,
		Message: "ads",
		Data:    nil,
	})
}

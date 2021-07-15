package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/global"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
	"github.com/saicem/todo/model/todomodel"
	"net/http"
)

// TodoChange
// @Summary TodoChange
// @Description 修改一个todo
// @Router /vtodo/change [post]
// @Param todo_id query int true "todo的ID"
// @Param json body request.TodoItemReq true "修改后的todo"
// @Success 200 object response.Response
func TodoChange(c *gin.Context) {
	// 读取参数
	var todoItemReq request.TodoItemReq
	rawData, err1 := c.GetRawData()
	if err1 != nil {
		c.JSON(http.StatusOK, response.Response{
			Status:  response.ERROR,
			Message: "wrong json",
		})
		return
	}
	err2 := json.Unmarshal(rawData, &todoItemReq)
	if err2 != nil {
		c.JSON(http.StatusOK, response.Response{
			Status:  response.ERROR,
			Message: "wrong json",
		})
		return
	}
	todoId := c.Query("todo_id")
	// 从cookie 获取 userId
	cookie, _ := c.Cookie("SESSIONID")
	userId := middleware.GetUserIdFromCookie(cookie)
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
	todoItem.Importance = todoItemReq.Importance
	todoItem.Content = todoItemReq.Content
	db.Save(&todoItem)
	c.JSON(http.StatusOK, response.Response{
		Status:  response.OK,
		Message: "修改一个todo",
		Data:    nil,
	})
}

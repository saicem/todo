package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoUpdate
// @Summary TodoUpdate
// @Description 修改一个todo
// @Router /vtodo/update [post]
// @Param todo_id query int true "todo的ID"
// @Param json body request.TodoItemReq true "修改后的todo"
// @Success 200 object response.Response
func TodoUpdate(c *gin.Context) {
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
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	// 查询数据库
	isSuccess := db.UpdateTodoItem(userId, todoItemReq, todoId)
	if isSuccess == false {
		c.JSON(http.StatusOK, response.Response{
			Status:  response.ERROR,
			Message: "error",
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Status:  response.OK,
		Message: "修改一个todo",
		Data:    nil,
	})
}

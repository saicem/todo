package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoPut
// @Summary TodoPut
// @Description 修改一个todo
// @Router /todo/{id} [put]
// @Param id path int true "User ID"
// @Param json body request.TodoItemReq2 true "修改后的todo"
// @Success 200 object response.Response
func TodoPut(c *gin.Context) {
	var todoItemReq request.TodoItemReq2
	if err := c.ShouldBindJSON(&todoItemReq); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	todoId := c.Param("id")
	cookie, _ := c.Cookie("SESSIONID")
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	// 查询数据库
	isSuccess := db.UpdateTodoItem(userId, todoItemReq, todoId)
	if isSuccess == false {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
	c.JSON(http.StatusOK, response.Response{
		Msg: "修改一个todo",
	})
}
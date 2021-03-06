package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
)

// TodoAdd
// @Summary TodoAdd
// @Description 添加一个新的todo
// @Param json body request.TodoItemReq1 true "新增的todo"
// @Router /todo/add [post]
// @Success 200 object response.Response
func TodoAdd(c *gin.Context) {
	userId, _ := c.Get("USERID")
	var todoItemReq *request.TodoItemReq1
	if err := c.BindJSON(&todoItemReq); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// todo 判断数据是否有效 是否可以存储
	isSuccess := db.CreateTodoItem(userId.(int), todoItemReq)
	if isSuccess {
		c.JSON(http.StatusOK, response.Response{
			Msg: "success",
		})
		return
	}
	c.JSON(http.StatusNotAcceptable, response.Response{
		Msg: "wrong group id",
	})
}

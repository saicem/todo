package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoAdd
// @Summary TodoAdd
// @Description 添加一个新的todo
// @Param json body request.TodoItemReq1 true "新增的todo"
// @Router /todo/add [post]
// @Success 200 object response.Response
func TodoAdd(c *gin.Context) {
	// 从cookie 获取 userId
	// todo 前边已经已经使用 sessionid 读取过了 userid 这里又读取了一遍 如何解决
	cookie, _ := c.Cookie("SESSIONID")
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	var todoItemReq request.TodoItemReq1
	if err := c.BindJSON(&todoItemReq); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// todo 判断数据是否有效 是否可以存储
	db.CreateTodoItem(userId, todoItemReq)
	c.JSON(http.StatusOK, response.Response{
		Msg: "success",
	})
}

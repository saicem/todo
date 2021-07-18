package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
)

// TodoGroupAdd
// @Summary TodoGroupAdd
// @Description 添加一个新的todoGroup
// @Param json body request.TodoGroupReq true "新增的todoGroup"
// @Router /todo_group/add [post]
// @Success 200 object response.Response
func TodoGroupAdd(c *gin.Context) {
	userId, _ := c.Get("userId")
	var todoGroupReq request.TodoGroupReq
	if err := c.BindJSON(&todoGroupReq); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	db.CreateTodoGroup(userId.(int), todoGroupReq)
	c.JSON(http.StatusOK, response.Response{
		Msg: "success",
	})
}

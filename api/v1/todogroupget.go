package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/model/response"
)

// TodoGroupGet
// @Summary TodoGroupGet
// @Description 查看已有的TodoGroup
// @Router /todo_group/list [get]
// @Success 200 object response.Response
func TodoGroupGet(c *gin.Context) {
	userId, _ := c.Get("userId")
	todoGroups := db.GetTodoGroups(userId.(int))
	c.JSON(http.StatusOK, response.Response{
		Msg:  "success",
		Data: todoGroups,
	})
}

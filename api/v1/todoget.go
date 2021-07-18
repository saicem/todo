package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/model/response"
)

// TodoGet
// @Summary TodoGet
// @Description 查看已有的todo
// @Router /todo/list [get]
// @Success 200 object response.Response
func TodoGet(c *gin.Context) {
	userId, _ := c.Get("userId")
	todoItems := db.GetTodoItems(userId.(int))
	c.JSON(http.StatusOK, response.Response{
		Msg:  "success",
		Data: todoItems,
	})
}

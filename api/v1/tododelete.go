package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/model/response"
)

// TodoDelete
// @Summary TodoDelete
// @Description 删除一个todo
// @Router /todo/{id} [delete]
// @Param id path int true "TodoId"
// @Success 200 object response.Response
func TodoDelete(c *gin.Context) {
	todoId := c.Param("id")
	userId, _ := c.Get("userId")
	isSuccess := db.DeleteTodoItem(userId.(int), todoId)
	if !isSuccess {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
	c.JSON(http.StatusOK, response.Response{
		Msg: "success",
	})
}

package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/model/response"
)

// TodoGroupDelete
// @Summary TodoGroupDelete
// @Description 删除一个 TodoGroup
// @Router /todo_group/{id} [delete]
// @Param id path int true "TodoGroupID"
// @Success 200 object response.Response
func TodoGroupDelete(c *gin.Context) {
	todoGroupId := c.Param("id")
	userId, _ := c.Get("USERID")
	isSuccess := db.DeleteTodoGroup(userId.(int), todoGroupId)
	if !isSuccess {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Msg: "success",
	})
}

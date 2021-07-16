package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoGroupDelete
// @Summary TodoGroupDelete
// @Description 删除一个 TodoGroup
// @Router /todo_group/{id} [delete]
// @Param id path int true "TodoGroupID"
// @Success 200 object response.Response
func TodoGroupDelete(c *gin.Context) {
	todoGroupId := c.Param("id")
	cookie, _ := c.Cookie("SESSIONID")
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	isSuccess := db.DeleteTodoGroup(userId, todoGroupId)
	if !isSuccess {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Msg: "success",
	})
}

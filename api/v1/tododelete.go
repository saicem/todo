package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoDelete
// @Summary TodoDelete
// @Description 删除一个todo
// @Router /todo/{id} [delete]
// @Param id path int true "TodoId"
// @Success 200 object response.Response
func TodoDelete(c *gin.Context) {
	todoId := c.Param("id")
	cookie, _ := c.Cookie("SESSIONID")
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	isSuccess := db.DeleteTodoItem(userId, todoId)
	if !isSuccess {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
	c.JSON(http.StatusOK, response.Response{
		Msg: "success",
	})
}

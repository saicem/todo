package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoGroupPut
// @Summary TodoGroupPut
// @Description
// @Router /todo_group/{id} [put]
// @Param id path int true "User Group Id"
// @Param json body request.TodoGroupReq true "修改后的todoGroup"
// @Success 200 object response.Response
func TodoGroupPut(c *gin.Context) {
	todoGroupId := c.Param("id")
	cookie, _ := c.Cookie("SESSIONID")
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	var todoGroupReq request.TodoGroupReq
	if err := c.ShouldBindJSON(&todoGroupReq); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	isSuccess := db.UpdateTodoGroup(userId, todoGroupReq, todoGroupId)
	if isSuccess == false {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}
	c.JSON(http.StatusOK, response.Response{Msg: "success"})
}

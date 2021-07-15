package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/global"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/response"
	"github.com/saicem/todo/model/todomodel"
	"net/http"
)

// TodoGet
// @Summary TodoGet
// @Description 查看已有的todo
// @Router /vtodo/get [get]
// @Success 200 object response.Response
func TodoGet(c *gin.Context) {
	db := global.Mysql
	var todoItems []todomodel.TodoItem
	// 从cookie 获取 userId
	cookie, _ := c.Cookie("SESSIONID")
	userId, _ := middleware.GetUserIdFromCookie(cookie)
	db.Where("uid = ?", userId).Find(&todoItems)
	c.JSON(http.StatusOK, response.Response{
		Status:  response.OK,
		Message: "success",
		Data:    todoItems,
	})
	return
}

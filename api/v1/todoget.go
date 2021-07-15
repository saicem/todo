package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/global"
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
	var todoItems todomodel.TodoItem
	// todo 判断是哪个用户
	db.Where("uid = ?", 1).Find(&todoItems)
	c.JSON(http.StatusOK, response.Response{
		Status:  response.OK,
		Message: "success",
		Data:    todoItems,
	})
	return
}

package v1

import (
	"github.com/gin-gonic/gin"
	. "github.com/saicem/todo/model/response"
	"net/http"
)

// TodoAdd
// @Summary TodoAdd
// @Description 添加一个新的todo
// @Router /vtodo/add [post]
// @Success 200 object response.Response
func TodoAdd(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:  OK,
		Message: "success",
		Data:    "this is a data",
	})
}

// TodoComplete
// @Summary TodoComplete
// @Description 一个todo完成了
// @Router /vtodo/complete [post]
// @Success 200 object response.Response
func TodoComplete(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:  OK,
		Message: "s",
		Data:    nil,
	})
}

// TodoChange
// @Summary TodoChange
// @Description 修改一个todo
// @Router /vtodo/change [post]
// @Success 200 object response.Response
func TodoChange(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:  OK,
		Message: "s",
		Data:    nil,
	})
}

// TodoDelete
// @Summary TodoDelete
// @Description 删除一个todo
// @Router /vtodo/delete [delete]
// @Success 200 object response.Response
func TodoDelete(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:  OK,
		Message: "ads",
		Data:    nil,
	})
}

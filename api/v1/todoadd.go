package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/middleware"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
	"net/http"
)

// TodoAdd
// @Summary TodoAdd
// @Description 添加一个新的todo
// @Param json body request.TodoItemReq true "新增的todo"
// @Router /vtodo/add [post]
// @Success 200 object response.Response
func TodoAdd(c *gin.Context) {
	// 获取参数
	var todoItemReq request.TodoItemReq
	rawData, err1 := c.GetRawData()
	if err1 != nil {
		c.JSON(http.StatusOK, response.Response{
			Status:  response.ERROR,
			Message: "wrong json",
		})
		return
	}
	err2 := json.Unmarshal(rawData, &todoItemReq)
	if err2 != nil {
		c.JSON(http.StatusOK, response.Response{
			Status:  response.ERROR,
			Message: "wrong json",
		})
		return
	}
	// 从cookie 获取 userId
	// todo 前边已经已经使用 sessionid 读取过了 userid 这里又读取了一遍 如何解决
	cookie, _ := c.Cookie("SESSIONID")
	uid, _ := middleware.GetUserIdFromCookie(cookie)
	// 存储数据
	db.CreateTodoItem(uid, todoItemReq)
	c.JSON(http.StatusOK, response.Response{
		Status:  response.OK,
		Message: "success",
		Data:    "this is a data",
	})
}

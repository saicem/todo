package v1

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saicem/todo/db"
	"github.com/saicem/todo/model/response"
)

// TodoGet
// @Summary TodoGet
// @Description 查看已有的todo
// @Router /todo/list [get]
// @Param create_at query string false "create_at"
// @Param keyword query string false "keyword"
// @Param todo_group_id query int false "todo_group_id"
// @Param is_finished query bool false "is_finished"
// @Success 200 object response.Response
func TodoGet(c *gin.Context) {
	userId, _ := c.Get("USERID")
	createAt, isGetCreateAt := c.GetQuery("create_at")
	keyword, isGetKeyword := c.GetQuery("keyword")
	todoGroupId, isGetTodoGroupId := c.GetQuery("todo_group_id")
	isFinished, isGetIsFinish := c.GetQuery("is_finished")
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if isGetCreateAt {
		if _, err := time.ParseInLocation("2006-01-02 15:04:05", createAt, loc); err != nil {
			log.Println(err)
			c.JSON(http.StatusNotAcceptable, response.Response{
				Msg: "wrong time format",
			})
			return
		}
	}
	if isGetTodoGroupId {
		if _, err := strconv.ParseInt(todoGroupId, 10, 64); err != nil {
			log.Println(err)
			c.JSON(http.StatusNotAcceptable, response.Response{
				Msg: "wrong group id",
			})
			return
		}
	}
	if isGetIsFinish {
		if _, err := strconv.ParseBool(isFinished); err != nil {
			log.Println(err)
			c.JSON(http.StatusNotAcceptable, response.Response{
				Msg: "wrong type of is_finish",
			})
			return
		}
	}
	todoItems := db.GetTodoItems(userId.(int), createAt, isGetCreateAt, keyword, isGetKeyword, todoGroupId, isGetTodoGroupId, isFinished, isGetIsFinish)
	c.JSON(http.StatusOK, response.Response{
		Msg:  "success",
		Data: todoItems,
	})
}

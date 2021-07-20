package db

import (
	"github.com/saicem/todo/global"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/todomodel"
)

func GetTodoItems(
	userId int,
	createAt string,
	isGetCreateAt bool,
	keyword string,
	isGetKeyword bool,
	todoGroupId string,
	isGetTodoGroupId bool,
	isFinish string,
	isGetIsFinish bool) *[]todomodel.TodoItem {
	var todoItems []todomodel.TodoItem
	sql := global.Mysql.Where("user_id = ?", userId)
	if isGetCreateAt {
		sql = sql.Where("create_at < ?", createAt)
	}
	if isGetKeyword {
		sql = sql.Where("LIKE", "%"+keyword+"%")
	}
	if isGetTodoGroupId {
		sql = sql.Where("group_id =", todoGroupId)
	}
	if isGetIsFinish {
		sql = sql.Where("is_finish = ", isFinish)
	}
	sql.Find(&todoItems)
	return &todoItems
}

func CreateTodoItem(userId int, req *request.TodoItemReq1) bool {
	res := global.Mysql.Create(&todomodel.TodoItem{
		TodoTitle:   req.TodoTitle,
		TodoContent: req.TodoContent,
		TodoGroupId: req.TodoGroupId,
		UserId:      userId,
	})
	return res.Error == nil
}

func UpdateTodoItem(uid int, req interface{}, todoId string) bool {
	todoItem, isSuccess := findTodoItem(uid, todoId)
	if !isSuccess {
		return false
	}
	res := global.Mysql.Model(&todoItem).Omit("id").Updates(req)
	if res.Error != nil {
		return false
	}
	return true
}

func DeleteTodoItem(userId int, todoId string) bool {
	todoItem, isSuccess := findTodoItem(userId, todoId)
	if !isSuccess {
		return false
	}
	global.Mysql.Delete(&todoItem)
	return true
}

func findTodoItem(userId int, todoId string) (*todomodel.TodoItem, bool) {
	var todoItem *todomodel.TodoItem
	res := global.Mysql.Find(&todoItem, "user_id = ? AND todo_id = ?", userId, todoId)
	if res.Error != nil {
		return nil, false
	}
	return todoItem, true
}

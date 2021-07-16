package db

import (
	"github.com/saicem/todo/global"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/response"
	"github.com/saicem/todo/model/todomodel"
)

func GetTodoItems(uid uint) *[]response.TodoItemRes {
	var todoItems []todomodel.TodoItem
	global.Mysql.Where("uid = ?", uid).Find(&todoItems)
	var todoItemsRes = make([]response.TodoItemRes, len(todoItems))
	for i := 0; i < len(todoItems); i++ {
		todoItemsRes[i] = *new(response.TodoItemRes).New(&todoItems[i])
	}

	return &todoItemsRes
}

func CreateTodoItem(uid uint, req request.TodoItemReq) bool {
	global.Mysql.Create(&todomodel.TodoItem{
		Importance: req.Importance,
		Content:    req.Content,
		Tag:        req.Tag,
		Uid:        uid,
	})
	return true
}

func findTodoItem(uid uint, todoId string) *todomodel.TodoItem {
	var todoItem todomodel.TodoItem
	global.Mysql.Find(&todoItem, "uid = ? AND id = ?", uid, todoId)
	return &todoItem
}

func UpdateTodoItem(uid uint, req request.TodoItemReq, todoId string) bool {
	todoItem := findTodoItem(uid, todoId)
	if todoItem == nil {
		return false
	}
	todoItem.Importance = req.Importance
	todoItem.Content = req.Content
	global.Mysql.Save(&todoItem)
	return true
}

func CompleteTodoItem(uid uint, todoId string) bool {
	todoItem := findTodoItem(uid, todoId)
	if todoItem == nil {
		return false
	}
	todoItem.Complete = true
	global.Mysql.Save(&todoItem)
	return true
}

func DeleteTodoItem(uid uint, todoId string) bool {
	todoItem := findTodoItem(uid, todoId)
	if todoItem == nil {
		return false
	}
	global.Mysql.Delete(&todoItem)
	return true
}

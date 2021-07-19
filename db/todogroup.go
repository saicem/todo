package db

import (
	"log"

	"github.com/saicem/todo/global"
	"github.com/saicem/todo/model/request"
	"github.com/saicem/todo/model/todomodel"
)

func GetTodoGroups(userId int) *[]todomodel.TodoGroup {
	var todoGroups []todomodel.TodoGroup
	global.Mysql.Where("user_id = ?", userId).Find(&todoGroups)
	return &todoGroups
}

func CreateTodoGroup(userId int, req request.TodoGroupReq) bool {
	res := global.Mysql.Create(&todomodel.TodoGroup{
		TodoGroupName: req.TodoGroupName,
		UserId:        userId,
	})
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	return true
}

func UpdateTodoGroup(userId int, req request.TodoGroupReq, todoGroupId string) bool {
	todoGroup := findTodoGroup(userId, todoGroupId)
	if todoGroup == nil {
		return false
	}
	todoGroup.TodoGroupName = req.TodoGroupName
	global.Mysql.Save(&todoGroup)
	return true
}

func DeleteTodoGroup(userId int, todoGroupId string) bool {
	todoGroup := findTodoGroup(userId, todoGroupId)
	if todoGroup == nil {
		return false
	}
	global.Mysql.Delete(&todoGroup)
	return true
}

func findTodoGroup(userId int, todoGroupId string) *todomodel.TodoGroup {
	var todoGroup todomodel.TodoGroup
	global.Mysql.Find(&todoGroup, "user_id = ? AND todo_group_id = ?", userId, todoGroupId)
	return &todoGroup
}

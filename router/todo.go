package router

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/hellojukay/todo-server/models"
)

func TodoList(w http.ResponseWriter, r *http.Request) {
	var todo = models.ToDo{
		Title:       "一个事情",
		Description: "这里是描述信息，非常的长",
	}
	var todoList []models.ToDo
	for i := 0; i < 100; i++ {
		todo.ID = int64(i)
		todoList = append(todoList, todo)
	}
	render.JSON(w, r, todoList)
}

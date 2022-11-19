package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/render"
	"github.com/hellojukay/todo-server/models"
)

func TodoList(w http.ResponseWriter, r *http.Request) {
	_finished := r.URL.Query().Get("finished")
	finished, _ := strconv.ParseBool(_finished)

	var task = models.ListAllTasks(finished)
	render.JSON(w, r, task)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var title = r.FormValue("title")
	log.Printf("新任务: %s", title)
	task, err := models.SaveTask(models.Task{
		Title: title,
	})
	if err != nil {
		log.Printf("无法创建任务 %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, task)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var _id = chi.URLParam(r, "taskID")
	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		log.Printf("非法任务编号 %s %s ", _id, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = models.RemoveTask(id); err != nil {
		log.Printf("无法删除任务 %d,数据库错误  %s ", id, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var _id = chi.URLParam(r, "taskID")
	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		log.Printf("非法任务编号 %s %s ", _id, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	log.Print("更新任务 ", task.Title, task.Finished)
	task.ID = id
	if err = models.UpdateTask(task); err != nil {
		log.Printf("无法更新任务 %d,数据库错误  %s ", id, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

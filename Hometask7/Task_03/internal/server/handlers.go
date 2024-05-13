package server

import (
	"Task_02/internal/data"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := data.GetTasksJson("assets/tasks.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index.html", tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func completeTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := data.GetTasksJson("assets/tasks.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sid := r.PathValue("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	taskId, err := data.GetTaskNumById(tasks, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasks[taskId].IsCompleted = true
	err = data.SaveTasksJson("assets/tasks.json", tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "row", tasks[taskId])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

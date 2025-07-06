package api

import (
	"net/http"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, map[string]string{"error": "Failed to retrieve tasks"})
		return
	}
	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}

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
		writeError(w, http.StatusInternalServerError, "Error fetching tasks")
		return
	}
	writeJson(w, http.StatusOK, tasksResponse{
		Tasks: tasks,
	})
}

package api

import (
	"fmt"
	"net/http"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	tasks, err := db.Tasks(50, search)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching tasks: %v", err))
		return
	}
	writeJson(w, http.StatusOK, tasksResponse{
		Tasks: tasks,
	})
}

package api

import (
	"fmt"
	"net/http"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		writeError(w, http.StatusBadRequest, "Task id is required")
		return
	}

	task, err := db.GetTask(id)

	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching task: %v", err))
		return
	}

	writeJson(w, http.StatusOK, task)
}

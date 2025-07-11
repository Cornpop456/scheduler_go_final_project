package api

import (
	"errors"
	"net/http"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if err := db.DeleteTask(id); err != nil {
		if errors.Is(err, db.ErrWrongId) {
			writeError(w, http.StatusNotFound, "Task not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	writeJson(w, http.StatusOK, emptyResponse{})
}

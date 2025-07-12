package api

import (
	"net/http"
	"time"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

func completeTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		writeError(w, http.StatusBadRequest, "Task id is required")
		return
	}

	task, err := db.GetTask(id)

	if err != nil {
		if err == db.ErrWrongId {
			writeError(w, http.StatusNotFound, "Task not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Error fetching task")
		return
	}

	if task.Repeat == "" {
		err := db.DeleteTask(task.ID)

		if err != nil {
			writeError(w, http.StatusInternalServerError, "Error deleting task")
			return
		}

		writeJson(w, http.StatusOK, emptyResponse{})
		return
	}

	nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error calculating next date")
		return
	}

	err = db.UpdateDate(nextDate, task.ID)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error updating task date")
		return
	}

	writeJson(w, http.StatusOK, emptyResponse{})
}

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Error reading request body: %v", err))
		return
	}

	var task db.Task

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if task.Title == "" || strings.TrimSpace(task.Title) == "" {
		writeError(w, http.StatusBadRequest, "Title is required")
		return
	}

	nowDate := time.Now()
	nextDate := ""

	if task.Date == "" {
		task.Date = nowDate.Format(db.TimeLayout)
	}

	if task.Repeat != "" {
		nextDate, err = NextDate(time.Now(), task.Date, task.Repeat)

		if err != nil {
			writeError(w, http.StatusBadRequest, fmt.Sprintf("Invalid repeat format: %v", err))
			return
		}
	}

	taskDate, err := time.Parse(db.TimeLayout, task.Date)

	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Invalid date format: %v", err))
		return
	}

	if afterNow(nowDate, taskDate) {
		task.Date = nextDate

		if task.Date == "" {
			task.Date = nowDate.Format(db.TimeLayout)
		}
	}

	err = db.UpdateTask(&task)

	if err != nil {
		if errors.Is(err, db.ErrWrongId) {
			writeError(w, http.StatusBadRequest, "Task not found")
			return
		}
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating task: %v", err))
		return
	}

	writeJson(w, http.StatusOK, emptyResponse{})
}

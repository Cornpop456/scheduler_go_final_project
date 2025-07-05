package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

type OkResponse struct {
	ID int64 `json:"id"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, ErrorResponse{ErrorMessage: fmt.Sprintf("Error reading request body: %v", err)})
		return
	}

	var task db.Task

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, ErrorResponse{ErrorMessage: fmt.Sprintf("Error parsing JSON: %v", err)})
		return
	}

	if task.Title == "" || strings.TrimSpace(task.Title) == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, ErrorResponse{ErrorMessage: "Title is required"})
		return
	}

	nowDate := time.Now()
	nextDate := ""

	if task.Date == "" {
		task.Date = nowDate.Format(timeLayout)
	}

	if task.Repeat != "" {
		nextDate, err = NextDate(time.Now(), task.Date, task.Repeat)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeJson(w, ErrorResponse{ErrorMessage: fmt.Sprintf("Invalid repeat format: %v", err)})
			return
		}
	}

	taskDate, err := time.Parse(timeLayout, task.Date)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, ErrorResponse{ErrorMessage: fmt.Sprintf("Invalid date format: %v", err)})
		return
	}

	if afterNow(nowDate, taskDate) {
		if nextDate == "" {
			task.Date = nowDate.Format(timeLayout)
		} else {
			task.Date = nextDate
		}
	}

	id, err := db.AddTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, ErrorResponse{ErrorMessage: fmt.Sprintf("Error adding task: %v", err)})
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeJson(w, OkResponse{ID: id})
}

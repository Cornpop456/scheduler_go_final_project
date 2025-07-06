package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

type tasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

type addTaskResponse struct {
	ID int64 `json:"id"`
}

type errorResponse struct {
	ErrorMessage string `json:"error"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
	}
}
func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(errorResponse{ErrorMessage: message}); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
	}
}

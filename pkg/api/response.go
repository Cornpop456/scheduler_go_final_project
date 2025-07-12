package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

type tasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

type addTaskResponse struct {
	ID int64 `json:"id"`
}

type authSuccessResponse struct {
	Token string `json:"token"`
}

type errorResponse struct {
	ErrorMessage string `json:"error"`
}

type emptyResponse struct{}

func writeJson(w http.ResponseWriter, status int, data any) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	_, err := w.Write(buf.Bytes())

	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJson(w, status, errorResponse{ErrorMessage: message})
}

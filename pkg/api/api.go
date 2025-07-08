package api

import "net/http"

func Init(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/nextdate", nextDayHandler)

	mux.HandleFunc("POST /api/task", addTaskHandler)
	mux.HandleFunc("GET /api/task", getTaskHandler)
	mux.HandleFunc("PUT /api/task", updateTaskHandler)
	mux.HandleFunc("DELETE /api/task", deleteTaskHandler)

	mux.HandleFunc("POST /api/task/done", completeTaskHandler)

	mux.HandleFunc("GET /api/tasks", tasksHandler)
}

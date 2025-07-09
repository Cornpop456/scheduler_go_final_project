package api

import "net/http"

func Init(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/nextdate", nextDayHandler)

	mux.HandleFunc("POST /api/task", auth(addTaskHandler))
	mux.HandleFunc("GET /api/task", auth(getTaskHandler))
	mux.HandleFunc("PUT /api/task", auth(updateTaskHandler))
	mux.HandleFunc("DELETE /api/task", auth(deleteTaskHandler))

	mux.HandleFunc("POST /api/task/done", auth(completeTaskHandler))

	mux.HandleFunc("GET /api/tasks", auth(tasksHandler))

	mux.HandleFunc("POST /api/signin", authHandler)
}

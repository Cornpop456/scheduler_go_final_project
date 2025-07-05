package api

import "net/http"

func Init() {
	http.HandleFunc("GET /api/nextdate", nextDayHandler)
	http.HandleFunc("POST /api/task", addTaskHandler)
}

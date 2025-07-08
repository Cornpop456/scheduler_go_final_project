package server

import (
	"net/http"
	"os"
	"time"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/api"
)

const webDir = "web"

var port = ":7540"

func New() *http.Server {
	if os.Getenv("TODO_PORT") != "" {
		port = ":" + os.Getenv("TODO_PORT")
	}

	mux := http.NewServeMux()

	api.Init(mux)

	mux.Handle("GET /", http.FileServer(http.Dir(webDir)))

	return &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

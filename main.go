package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/api"
	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
)

const webDir = "web"

var port = ":7540"
var dbFile = "scheduler.db"

func main() {
	if os.Getenv("TODO_PORT") != "" {
		port = ":" + os.Getenv("TODO_PORT")
	}

	if os.Getenv("TODO_DBFILE") != "" {
		dbFile = os.Getenv("TODO_DBFILE")
	}

	err := db.Init(dbFile)

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	api.Init()

	http.Handle("GET /", http.FileServer(http.Dir(webDir)))

	log.Println("Server running on http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}

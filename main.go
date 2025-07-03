package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/api"
	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
	"github.com/Cornpop456/scheduler_go_final_project/tests"
)

const webDir = "web"

var port = ":" + strconv.Itoa(tests.Port)
var dbFile = tests.DBFile

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

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	log.Println("Server running on http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}

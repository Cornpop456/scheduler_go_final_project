package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
	"github.com/Cornpop456/scheduler_go_final_project/tests"
)

const WEB_DIR = "web"

var PORT = ":" + strconv.Itoa(tests.Port)
var DB_FILE = tests.DBFile

func main() {
	if os.Getenv("TODO_PORT") != "" {
		PORT = ":" + os.Getenv("TODO_PORT")
	}

	if os.Getenv("TODO_DBFILE") != "" {
		DB_FILE = os.Getenv("TODO_DBFILE")
	}

	err := db.Init(DB_FILE)

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	http.Handle("/", http.FileServer(http.Dir(WEB_DIR)))

	log.Println("Server running on http://localhost" + PORT)

	log.Fatal(http.ListenAndServe(PORT, nil))
}

package main

import (
	"log"
	"os"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
	"github.com/Cornpop456/scheduler_go_final_project/pkg/server"
)

var dbFile = "scheduler.db"

func main() {
	if os.Getenv("TODO_DBFILE") != "" {
		dbFile = os.Getenv("TODO_DBFILE")
	}

	err := db.Init(dbFile)

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	server := server.New()

	log.Println("Server running on http://localhost" + server.Addr)

	log.Fatal(server.ListenAndServe())
}

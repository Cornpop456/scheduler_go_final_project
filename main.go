package main

import (
	"log"
	"os"

	"github.com/Cornpop456/scheduler_go_final_project/pkg/config"
	"github.com/Cornpop456/scheduler_go_final_project/pkg/db"
	"github.com/Cornpop456/scheduler_go_final_project/pkg/server"
)

func main() {
	dbFileEnv := os.Getenv("TODO_DBFILE")
	passwordEnv := os.Getenv("TODO_PASSWORD")
	portEnv := os.Getenv("TODO_PORT")

	conf := config.New(passwordEnv, portEnv, dbFileEnv)

	err := db.Init(conf.DbFile)

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	defer db.Close()

	server := server.New(conf.Port, conf.Password)

	log.Println("Server running on http://localhost" + server.Addr)

	log.Fatal(server.ListenAndServe())
}

package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Cornpop456/scheduler_go_final_project/tests"
)

const WEB_DIR = "web"

var PORT = ":" + strconv.Itoa(tests.Port)

func main() {
	if os.Getenv("TODO_PORT") != "" {
		PORT = ":" + os.Getenv("TODO_PORT")
	}

	http.Handle("/", http.FileServer(http.Dir(WEB_DIR)))

	log.Println("Server running on http://localhost" + PORT)

	log.Fatal(http.ListenAndServe(PORT, nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ASHWINK07/tasker/Controller"
	"github.com/ASHWINK07/tasker/Testing"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/testing", Testing.Testing)
	http.HandleFunc("/records/", Controller.Routing)
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}

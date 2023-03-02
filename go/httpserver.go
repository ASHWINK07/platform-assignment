package main

/*

This is the main file which runs the http server
*/
import (
	"fmt"
	"log"
	"net/http"

	"github.com/ASHWINK07/tasker/Controller"
	"github.com/ASHWINK07/tasker/Testing"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// testing is just to check whether the server is running or not
	http.HandleFunc("/testing", Testing.Testing)
	//all the api request with path as /records are routed to Controller module
	http.HandleFunc("/records/", Controller.Routing)
	fmt.Printf("Starting server at port 8081\n")
	//http server running on port 8081
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}

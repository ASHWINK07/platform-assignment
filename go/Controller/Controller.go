package Controller

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ASHWINK07/tasker/MongoConnections"
	"github.com/ASHWINK07/tasker/MongoOperations"
	"github.com/ASHWINK07/tasker/SqlConnections"
	"github.com/ASHWINK07/tasker/SqlServices"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

func Routing(w http.ResponseWriter, r *http.Request) {
	// if the db query parameter i mongodb like - http://127.0.0.1:8081/records/db=mongodb
	if r.URL.Query().Get("db") == "mongodb" {
		//"mongodb://localhost:27017"
		client, ctx, cancel, err := MongoConnections.Connect()
		if err != nil {
			panic(err)
		}

		defer MongoConnections.Close(client, ctx, cancel)
		//get the id of the employee http://127.0.0.1:8081/records/1/db=mongodb  here id will return [records 1]
		id := strings.Split(r.URL.Path, "/")
		//if it is a Get Request on mongodb database
		if r.Method == "GET" {
			var results []bson.D
			//sample get request:curl http://127.0.0.1:8081/records/3/\?db\=mongodb
			//here get the details of employee with id as 3
			results, err := MongoOperations.MongoGet(id[2], client, ctx)
			if err != nil {
				panic(err)
				io.WriteString(w, "Id not found")
				fmt.Fprintf(w, "ID not found")
				return
			}
			var count int = 0
			for _, doc := range results {
				count = count + 1
				fmt.Println(doc)
			}
			if count == 0 {
				io.WriteString(w, "Id not found")
				return
			}
			var value1, value2 string

			value1 = results[0][0].Value.(string)
			value2 = results[0][1].Value.(string)
			fmt.Println(value1, value2)
			io.WriteString(w, value1+" "+value2)
			return
		} else if r.Method == "POST" {
			//if it is a Post Request on mongodb database
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)

				return
			}
			//get the form values i.e name and department
			//sample Post request :curl -X POST http://127.0.0.1:8081/records/\?db\=mongodb -d "name=ashwin&department=platform"
			name := r.FormValue("name")
			department := r.FormValue("department")
			//Insert into mongodb database
			err = MongoOperations.MongoInsert(name, department, client, ctx)
			//check whether there is error
			if err != nil {
				panic(err)
				io.WriteString(w, "Insertion failed")
				return
			}
			fmt.Println("POST request successfull")
			//return the ouput for Post request
			io.WriteString(w, "200 Post request Successfull")
			return
		} else if r.Method == "DELETE" {
			//if it is a Delete Request on mongodb database
			//sample delete request:curl -X DELETE http://127.0.0.1:8081/records/1/\?db\=mongodb
			//Delete the employee with id as 1
			err = MongoOperations.MongoDelete(id[2], client, ctx)
			if err != nil {
				panic(err)
				io.WriteString(w, "Id not found")
				fmt.Fprintf(w, "ID not found")
				return
			}
			fmt.Println("Deletion Successfull")
			io.WriteString(w, "200 Delete request Successfull")
			return
		} else if r.Method == "PUT" {
			//if it is a Put Request on mongodb database
			//Get the field values that is name and department
			//sample put request:curl -X PUT http://127.0.0.1:8081/records/2/\?db\=mongodb -d "name=ashwin&department=frontend"
			//update the employee with id as 2
			name := r.FormValue("name")
			department := r.FormValue("department")
			err = MongoOperations.MongoUpdate(id[2], name, department, client, ctx)
			if err != nil {
				panic(err)
				io.WriteString(w, "Id not found")
				fmt.Fprintf(w, "ID not found")
				return
			}
			fmt.Println("update single document")
			//return the ouput for Put request
			io.WriteString(w, "200 Put request Successfull")
			return

		}
	} else if r.URL.Query().Get("db") == "mysql" {
		//if the db query parameter is mysql
		//Connect to mysql database
		db, err := SqlConnections.SqlConnect()
		//retriev the employee id
		id := strings.Split(r.URL.Path, "/")
		if r.Method == "GET" {
			//sample get request:curl http://127.0.0.1:8081/records/3/\?db\=mysql
			//get details of employee with id as 3
			var username string
			var team string
			var userid int
			username, team, userid = SqlServices.SqlGet(db, id[2])
			//return the output
			if userid == -1 {
				io.WriteString(w, "Id not found")
				return
			}
			fmt.Fprintf(w, "_id = %d\n", userid)
			fmt.Fprintf(w, "Name = %s\n", username)
			fmt.Fprintf(w, "Department = %s\n", team)
			return

		} else if r.Method == "POST" {
			//sample post request:curl -X POST http://127.0.0.1:8081/records/\?db\=mysql -d "name=ashwin&department=platform"
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			name := r.FormValue("name")
			department := r.FormValue("department")
			err = SqlServices.SqlInsert(db, name, department)
			if err != nil {
				panic(err)
			} else {
				io.WriteString(w, "Mysql insertion successfull")
			}
			return
		} else if r.Method == "PUT" {
			//sample put request:curl -X PUT http://127.0.0.1:8081/records/2/\?db\=mysql -d "name=ashwin&department=frontend"
			//update the employee with id as 2
			name := r.FormValue("name")
			department := r.FormValue("department")
			err = SqlServices.SqlUpdate(db, id[2], name, department)
			if err != nil {
				panic(err)
				return
			} else {
				io.WriteString(w, "Mysql updation successfull")
			}
			return
		} else if r.Method == "DELETE" {
			//sample delete request:curl -X DELETE http://127.0.0.1:8081/records/1/\?db\=mysql
			//delete the employee with id as 1
			err = SqlServices.SqlDelete(db, id[2])
			if err != nil {
				panic(err)
			} else {
				//return the ouput for Delete request
				io.WriteString(w, "Mysql deletion successfull")
			}
			return
		}
	}
}

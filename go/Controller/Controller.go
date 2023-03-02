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
	if r.URL.Query().Get("db") == "mongodb" {
		//"mongodb://localhost:27017"
		client, ctx, cancel, err := MongoConnections.Connect()
		if err != nil {
			panic(err)
		}
		var a int = 0
		if a == 1 {
			defer MongoConnections.Close(client, ctx, cancel)
		}
		//

		id := strings.Split(r.URL.Path, "/")
		//var document interface{}
		if r.Method == "GET" {
			// var filter, option interface{}
			// p, err := strconv.Atoi(id[2])
			// filter = bson.D{
			// 	{"_id", p},
			// }
			// option = bson.D{{"_id", 0}}
			// cursor, err := MongoServices.Query(client, ctx, "employee", "records", filter, option)
			// if err != nil {
			// 	panic(err)
			// }
			// var results []bson.D
			// if err := cursor.All(ctx, &results); err != nil {
			// 	panic(err)
			// }
			// fmt.Println(results[0][0])
			// fmt.Println(results[0][1])
			var results []bson.D
			results, err := MongoOperations.MongoGet(id[2], client, ctx)
			if err != nil {
				panic(err)
				return
			}

			fmt.Println(results)
			//io.WriteString(w,)

			io.WriteString(w, "200 Get request Successfull")
			return
		} else if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			name := r.FormValue("name")
			department := r.FormValue("department")
			err = MongoOperations.MongoInsert(name, department, client, ctx)
			if err != nil {
				panic(err)
				return
			}
			fmt.Println("POST request successfull")
			io.WriteString(w, "200 Post request Successfull")
			return
		} else if r.Method == "DELETE" {
			err = MongoOperations.MongoDelete(id[2], client, ctx)
			if err != nil {
				panic(err)
				return
			}
			fmt.Println("Deletion Successfull")
			io.WriteString(w, "200 Delete request Successfull")
			return
		} else if r.Method == "PUT" {
			// fmt.Println("Put Request Works")
			name := r.FormValue("name")
			department := r.FormValue("department")
			// fmt.Fprintf(w, "Name = %s\n", name)
			// fmt.Fprintf(w, "Department = %s\n", department)
			// tempid, _ := strconv.Atoi((id[2]))
			// filter := bson.D{
			// 	{"_id", tempid},
			// }
			// update := bson.D{
			// 	{"$set", bson.D{
			// 		{"Department", department},
			// 	}},
			// }
			err = MongoOperations.MongoUpdate(id[2], name, department, client, ctx)
			//result, err := MongoServices.UpdateOne(client, ctx, "employee", "records", filter, update)
			if err != nil {
				panic(err)
				return
			}
			fmt.Println("update single document")
			// fmt.Println(result.ModifiedCount)
			io.WriteString(w, "200 Put request Successfull")
			return

		}
	} else if r.URL.Query().Get("db") == "mysql" {
		db, err := SqlConnections.SqlConnect()
		id := strings.Split(r.URL.Path, "/")
		if r.Method == "GET" {
			var username string
			var team string
			var userid int
			username, team, userid = SqlServices.SqlGet(db, id[2])
			fmt.Fprintf(w, "_id = %d\n", userid)
			fmt.Fprintf(w, "Name = %s\n", username)
			fmt.Fprintf(w, "Department = %s\n", team)
			return

		} else if r.Method == "POST" {
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
			err = SqlServices.SqlDelete(db, id[2])
			if err != nil {
				panic(err)
			} else {
				io.WriteString(w, "Mysql deletion successfull")
			}
			return
		}
	}
}

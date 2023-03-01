package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	//"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return
}

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

func deleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {
	// select document and collection
	collection := client.Database(dataBase).Collection(col)
	// query is used to match a document  from the collection.
	result, err = collection.DeleteOne(ctx, query)
	return
}

func Testing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testing")
	return
}

// func mongodb(w http.ResponseWriter, r *http.Request) {
// 	client, ctx, cancel, err := connect("mongodb://localhost:27017")
// 	if err != nil {
// 		panic(err)
// 	}
// 	var a int = 0
// 	if a == 1 {
// 		defer close(client, ctx, cancel)
// 	}
// 	//
// 	var document interface{}
// 	if r.URL.Path != "/mongodb" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	} else if r.Method == "GET" {
// 		name := r.URL.Query().Get("name")
// 		fmt.Println("Name : =>", name)
// 		fmt.Println("Get Request Works")
// 		var filter, option interface{}
// 		filter = bson.D{
// 			{"Name", name},
// 		}
// 		option = bson.D{{"_id", 0}}
// 		cursor, err := query(client, ctx, "employee", "records", filter, option)
// 		if err != nil {
// 			panic(err)
// 		}
// 		var results []bson.D
// 		if err := cursor.All(ctx, &results); err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("Query Result")
// 		for _, doc := range results {
// 			fmt.Println(doc)
// 		}
// 		io.WriteString(w, "200 Get request Successfull")
// 	} else if r.Method == "POST" {
// 		fmt.Println("Post Request Works")
// 		if err := r.ParseForm(); err != nil {
// 			fmt.Fprintf(w, "ParseForm() err: %v", err)
// 			return
// 		}
// 		fmt.Fprintf(w, "POST request successful")
// 		name := r.FormValue("name")
// 		department := r.FormValue("department")
// 		fmt.Fprintf(w, "Name = %s\n", name)
// 		fmt.Fprintf(w, "Department = %s\n", department)

// 		document = bson.D{
// 			{"Name", name},
// 			{"Department", department},
// 		}
// 		insertOneResult, err := insertOne(client, ctx, "employee", "records", document)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Result of InsertOne")
// 		fmt.Println(insertOneResult.InsertedID)
// 		io.WriteString(w, "200 Post request Successfull")
// 	} else if r.Method == "PUT" {
// 		fmt.Println("Put Request Works")
// 		name := r.FormValue("name")
// 		department := r.FormValue("department")
// 		fmt.Fprintf(w, "Name = %s\n", name)
// 		fmt.Fprintf(w, "Department = %s\n", department)
// 		filter := bson.D{
// 			{"Name", name},
// 		}
// 		// The field of the document that need to updated.
// 		update := bson.D{
// 			{"$set", bson.D{
// 				{"Department", department},
// 			}},
// 		}
// 		result, err := UpdateOne(client, ctx, "employee", "records", filter, update)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// print count of documents that affected
// 		fmt.Println("update single document")
// 		fmt.Println(result.ModifiedCount)
// 		io.WriteString(w, "200 Put request Successfull")
// 	} else if r.Method == "DELETE" {
// 		name := r.URL.Query().Get("name")
// 		fmt.Println("name : =>", name)
// 		fmt.Println("Delete Request Works")
// 		query := bson.D{
// 			{"Name", name},
// 		}
// 		// Returns result of deletion and error
// 		result, err := deleteOne(client, ctx, "employee", "records", query)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// print the count of affected documents
// 		fmt.Println("No.of rows affected by DeleteOne()")
// 		fmt.Println(result.DeletedCount)
// 		io.WriteString(w, "200 Delete request Successfull")
// 	}
// 	fmt.Println("CRUD operation")

// }

// func mysql(w http.ResponseWriter, r *http.Request) {
// 	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/mysql")
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = db.Ping()

// 	// handle error
// 	if err != nil {
// 		panic(err)
// 	}
// 	//_, err = db.Exec("CREATE TABLE IF NOT EXISTS  users(Name text, Department text);")
// 	_, err = db.Exec("Create Table if not exists users(Name text, Department text,id bigint(20) unsigned NOT NULL AUTO_INCREMENT")
// 	fmt.Print("Pong\n")
// 	defer db.Close()
// 	fmt.Println("Mysql successfull")
// 	if r.Method == "GET" {
// 		name := r.URL.Query().Get("name")
// 		fmt.Println(name)
// 		var username string
// 		var team string
// 		selectquery := "Select * from users where name=?"
// 		err := db.QueryRow(selectquery, name).Scan(&username, &team)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Fprintf(w, "Name = %s\n", username)
// 		fmt.Fprintf(w, "Department = %s\n", team)

// 	} else if r.Method == "POST" {
// 		fmt.Println("Post Request Works")
// 		if err := r.ParseForm(); err != nil {
// 			fmt.Fprintf(w, "ParseForm() err: %v", err)
// 			return
// 		}
// 		//fmt.Fprintf(w, "POST request successful")
// 		name := r.FormValue("name")
// 		department := r.FormValue("department")
// 		fmt.Fprintf(w, "Name = %s\n", name)
// 		fmt.Fprintf(w, "Department = %s\n", department)
// 		insertquery := "INSERT INTO users(Name,Department) values(?,?)"
// 		_, err = db.ExecContext(context.Background(), insertquery, name, department)
// 		if err != nil {
// 			panic(err)
// 		} else {
// 			io.WriteString(w, "Mysql insertion successfull")
// 		}
// 	} else if r.Method == "PUT" {
// 		name := r.FormValue("name")
// 		department := r.FormValue("department")
// 		fmt.Fprintf(w, "Name = %s\n", name)
// 		fmt.Fprintf(w, "Department = %s\n", department)
// 		updatequery := "update users set Department=? where Name=?"
// 		_, err = db.ExecContext(context.Background(), updatequery, department, name)
// 		if err != nil {
// 			panic(err)
// 		} else {
// 			io.WriteString(w, "Mysql updation successfull")
// 		}
// 	} else if r.Method == "DELETE" {
// 		name := r.FormValue("name")
// 		deletequery := "DELETE from users where name=?"
// 		_, err = db.ExecContext(context.Background(), deletequery, name)
// 		if err != nil {
// 			panic(err)
// 		} else {
// 			io.WriteString(w, "Mysql deletion successfull")
// 		}
// 	}
// }

func routing(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("db") == "mongodb" {
		client, ctx, cancel, err := connect("mongodb://localhost:27017")
		if err != nil {
			panic(err)
		}
		var a int = 0
		if a == 1 {
			defer close(client, ctx, cancel)
		}
		//

		id := strings.Split(r.URL.Path, "/")
		var document interface{}
		if r.Method == "GET" {
			var filter, option interface{}
			//p, err := strconv.Atoi(id[2])
			filter = bson.D{
				{"_id", id[2]},
			}
			option = bson.D{{"_id", 0}}
			cursor, err := query(client, ctx, "employee", "records", filter, option)
			if err != nil {
				panic(err)
			}
			var results []bson.D
			if err := cursor.All(ctx, &results); err != nil {
				panic(err)
			}
			for _, doc := range results {
				fmt.Println(doc)
			}
			io.WriteString(w, "200 Get request Successfull")
		} else if r.Method == "POST" {
			fmt.Println("Post Request Works")
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			fmt.Fprintf(w, "POST request successful")
			name := r.FormValue("name")
			department := r.FormValue("department")
			//userid := r.FormValue("id")
			rand.Seed(time.Now().UnixNano())
			userid := rand.Intn(399)
			fmt.Fprintf(w, "Name = %s\n", name)
			fmt.Fprintf(w, "Department = %s\n", department)
			fmt.Fprintf(w, "user id=%s\n", userid)
			document = bson.D{
				{"Name", name},
				{"Department", department},
				{"_id", userid},
			}
			insertOneResult, err := insertOne(client, ctx, "employee", "records", document)
			if err != nil {
				panic(err)
			}
			fmt.Println("Result of InsertOne")
			fmt.Println(insertOneResult.InsertedID)
			io.WriteString(w, "200 Post request Successfull")
		} else if r.Method == "DELETE" {
			name := r.URL.Query().Get("name")
			fmt.Println("name : =>", name)
			fmt.Println("Delete Request Works")
			tempid, _ := strconv.Atoi(id[2])
			query := bson.D{
				{"_id", tempid},
			}
			result, err := deleteOne(client, ctx, "employee", "records", query)
			if err != nil {
				panic(err)
			}
			//fmt.Println("No.of rows affected by DeleteOne()")
			fmt.Println(result.DeletedCount)
			io.WriteString(w, "200 Delete request Successfull")
		} else if r.Method == "PUT" {
			fmt.Println("Put Request Works")
			name := r.FormValue("name")
			department := r.FormValue("department")
			fmt.Fprintf(w, "Name = %s\n", name)
			fmt.Fprintf(w, "Department = %s\n", department)
			tempid, _ := strconv.Atoi((id[2]))
			filter := bson.D{
				{"_id", tempid},
			}
			// The field of the document that need to updated.
			update := bson.D{
				{"$set", bson.D{
					{"Department", department},
				}},
			}
			result, err := UpdateOne(client, ctx, "employee", "records", filter, update)
			if err != nil {
				panic(err)
			}
			// print count of documents that affected
			fmt.Println("update single document")
			fmt.Println(result.ModifiedCount)
			io.WriteString(w, "200 Put request Successfull")

		}
	} else if r.URL.Query().Get("db") == "mysql" {
		db, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/mysql")
		if err != nil {
			panic(err)
		}
		err = db.Ping()

		// handle error
		if err != nil {
			panic(err)
		}
		//_, err = db.Exec("Create Table if not exists users(Name text, Department text,id bigint(20) unsigned NOT NULL AUTO_INCREMENT)")
		_, err = db.Exec("Create Table if not exists users(Name text, Department text,id INT PRIMARY KEY NOT NULL AUTO_INCREMENT)")
		defer db.Close()
		id := strings.Split(r.URL.Path, "/")
		if r.Method == "GET" {
			var username string
			var team string
			var userid int
			selectquery := "Select * from users where id=?"
			err := db.QueryRow(selectquery, id[2]).Scan(&username, &team, &userid)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "_id = %d\n", userid)
			fmt.Fprintf(w, "Name = %s\n", username)
			fmt.Fprintf(w, "Department = %s\n", team)

		} else if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			name := r.FormValue("name")
			department := r.FormValue("department")
			fmt.Fprintf(w, "Name = %s\n", name)
			fmt.Fprintf(w, "Department = %s\n", department)
			insertquery := "INSERT INTO users(Name,Department) values(?,?)"
			_, err = db.ExecContext(context.Background(), insertquery, name, department)
			if err != nil {
				panic(err)
			} else {
				io.WriteString(w, "Mysql insertion successfull")
			}
		} else if r.Method == "PUT" {
			name := r.FormValue("name")
			department := r.FormValue("department")
			fmt.Fprintf(w, "Name = %s\n", name)
			fmt.Fprintf(w, "Department = %s\n", department)
			updatequery := "update users set Department=? where id=?"
			p, err := strconv.Atoi(id[2])
			_, err = db.ExecContext(context.Background(), updatequery, department, p)
			if err != nil {
				panic(err)
			} else {
				io.WriteString(w, "Mysql updation successfull")
			}
		} else if r.Method == "DELETE" {
			p, err := strconv.Atoi(id[2])
			deletequery := "DELETE from users where id=?"
			_, err = db.ExecContext(context.Background(), deletequery, p)
			if err != nil {
				panic(err)
			} else {
				io.WriteString(w, "Mysql deletion successfull")
			}
		}
	}

}

func main() {
	http.HandleFunc("/testing", Testing)
	// http.HandleFunc("/mongodb", mongodb)
	// http.HandleFunc("/mysql", mysql)
	http.HandleFunc("/records/", routing)
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}

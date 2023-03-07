package SqlServices

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func SqlInsert(db *sql.DB, name, department string) error {
	//insert username and department ,id is auto-increamented need not be inserted
	insertquery := "INSERT INTO users(Department,Name) values(?,?)"
	_, err := db.ExecContext(context.Background(), insertquery, department, name)
	return err
}
func SqlGet(db *sql.DB, id string) (string, string, int) {
	var username string
	var team string
	var userid int
	//get the username,team and userid
	selectquery := "Select * from users where id=?"
	err := db.QueryRow(selectquery, id).Scan(&userid, &team, &username)
	if err != nil {
		fmt.Println("no details found")
		userid = -1
		team = "null"
		username = "null"
		return username, team, userid
		//panic(err)
	}
	return username, team, userid

}
func SqlUpdate(db *sql.DB, id, name, department string) error {
	//update the employee based on id
	updatequery := "update users set Department=?,Name=? where id=?"
	userid, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(context.Background(), updatequery, department, name, userid)
	return err
}

func SqlDelete(db *sql.DB, id string) error {
	//delete the emplopyee based on id
	//convert the id to int since the datatype declared in table is int
	userid, err := strconv.Atoi(id)
	if err != nil {
		panic((err))
	}
	deletequery := "DELETE from users where id=?"
	result, err := db.ExecContext(context.Background(), deletequery, userid)
	fmt.Println(*result)
	return err
}

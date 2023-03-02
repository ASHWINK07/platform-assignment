package SqlServices

import (
	"context"
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func SqlInsert(db *sql.DB, name, department string) error {
	//insert username and department ,id is auto-increamented need not be inserted
	insertquery := "INSERT INTO users(Name,Department) values(?,?)"
	_, err := db.ExecContext(context.Background(), insertquery, name, department)
	return err
}
func SqlGet(db *sql.DB, id string) (string, string, int) {
	var username string
	var team string
	var userid int
	//get the username,team and userid
	selectquery := "Select * from users where id=?"
	err := db.QueryRow(selectquery, id).Scan(&username, &team, &userid)
	if err != nil {
		panic(err)
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
	userid, err := strconv.Atoi(id)
	if err != nil {
		panic((err))
	}
	deletequery := "DELETE from users where id=?"
	_, err = db.ExecContext(context.Background(), deletequery, userid)
	return err
}

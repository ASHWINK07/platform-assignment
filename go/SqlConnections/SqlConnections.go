package SqlConnections

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host         = "localhost"
	port         = 3306
	user         = "root"
	password     = "12345"
	databasename = "mysql"
	//dbname = "postgres"
)

func SqlConnect() (*sql.DB, error) {
	//the connection url looks like :root:12345@tcp(localhost:3306)/mysql
	sqlconnectionurl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, databasename)
	db, err := sql.Open(databasename, sqlconnectionurl)
	_, err = db.Exec("Create Table if not exists users(Name text, Department text,id INT PRIMARY KEY NOT NULL AUTO_INCREMENT)")
	//return the pointer to database created
	return db, err
}

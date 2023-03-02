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
	sqlconnectionurl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, databasename)
	db, err := sql.Open(databasename, sqlconnectionurl)
	// create_table := "Create Table if not exists users(Name text, Department text,id INT PRIMARY KEY NOT NULL AUTO_INCREMENT)"
	// _, err = db.ExecContext(context.Background(), create_table)
	_, err = db.Exec("Create Table if not exists users(Name text, Department text,id INT PRIMARY KEY NOT NULL AUTO_INCREMENT)")

	return db, err
}

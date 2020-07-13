package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// if there is no such file, it will be created
	db, _ := sql.Open("sqlite3", "./mydb.db")

	// if there is no such table, it will be created with the following properties
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users ( id INTEGER PRIMARY KEY, nickname TEXT NOT NULL UNIQUE, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL )")
	statement.Exec()

	// to insert variables
	statement, _ = db.Prepare("INSERT INTO users (nickname, email, password) VALUES (?, ?, ?)")
	statement.Exec("Jane", "s@gmail.com", "12345678910") // exact these values will not work cause nickname "Jane" is already exists

	// to parse variables
	rows, _ := db.Query("SELECT id, nickname, email, password FROM users")
	var id int
	var nickname, email, password string
	for rows.Next() {
		rows.Scan(&id, &nickname, &email, &password)
		fmt.Printf("%d: %s %s %s\n", id, nickname, email, password)
	}
  //test
}

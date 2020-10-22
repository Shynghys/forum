package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	id       int
	username string
	email    string
	password string
	posts    Posts
	comments Comments
}

type Posts struct {
	id       int
	authorID int
	comments Comments
	likes    int
}

type Comments struct {
	id       int
	postID   int
	authorID int
	likes    int
}

func main() {

	// if there is no such file, it will be created
	db, _ := sql.Open("sqlite3", "./m.db")

	// if there is no such table, it will be created with the following properties
	statementForUsers, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users ( id INTEGER PRIMARY KEY, nickname TEXT NOT NULL UNIQUE, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL )")
	statementForUsers.Exec()
	// statementForPosts, _ := db.Prepare("CREATE TABLE IF NOT EXISTS posts ( id INTEGER PRIMARY KEY, authorID INTEGER PRIMARY KEY,  likes INTEGER )")
	// statementForPosts.Exec()
	// statementForComments, _ := db.Prepare("CREATE TABLE IF NOT EXISTS comments ( id INTEGER PRIMARY KEY, likes INTEGER )")
	// statementForComments.Exec()
	// to insert variables
	statementForUsers, _ = db.Prepare("INSERT INTO users (nickname, email, password) VALUES (?, ?, ?)")
	statementForUsers.Exec("ButerBrot359", "batowka359@gmail.com", "1234") // exact these values will not work cause nickname "Jane" is already exists

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

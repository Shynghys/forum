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
	created  string
	posts    Posts
	comments Comments
}

type Posts struct {
	id       int
	authorID int
	created  string
	comments Comments
	likes    int
}

type Comments struct {
	id       int
	postID   int
	authorID int
	created  string
	likes    int
}

func main() {
	fmt.Println("1")
	newDB := createDatabase()
	addUser(newDB, 1, "buter", "bat@mail.ru", "123", "1212")
	fmt.Println(newDB)
	rows, _ := newDB.Query("SELECT id, nickname, email, password FROM users")
	var id int
	var nickname, email, password string
	for rows.Next() {
		rows.Scan(&id, &nickname, &email, &password)
		fmt.Printf("%d: %s %s %s\n", id, nickname, email, password)
	}
}

func createDatabase() *sql.DB {
	db, _ := sql.Open("sqlite3", "./newDB.db")
	// checkErr(err)

	// if there is no such table, it will be created with the following properties
	statementForUsers, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users ( id INTEGER PRIMARY KEY, username TEXT NOT NULL UNIQUE, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, created TEXT )")
	statementForUsers.Exec()
	// statementForPosts, _ := db.Prepare("CREATE TABLE IF NOT EXISTS posts ( id INTEGER PRIMARY KEY, authorID INTEGER PRIMARY KEY, created TEXT, likes INTEGER )")
	// statementForPosts.Exec()
	// statementForComments, _ := db.Prepare("CREATE TABLE IF NOT EXISTS comments ( id INTEGER PRIMARY KEY,postID INTEGER, authorID INTEGER, created TEXT, likes INTEGER )")
	// statementForComments.Exec()
	return db
}

func addUser(db *sql.DB, id int, username string, email string, password string, created string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users (id, username, email, password, created) VALUES (?,?,?,?,?)")
	stmt.Exec(id, username, email, password, created)
	// _, err := stmt.Exec(id, username, email, password, created)
	// checkErr(err)
	tx.Commit()
}

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

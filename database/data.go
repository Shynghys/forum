package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/satori/go.uuid"
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
  	name string
	created  string
	comments Comments
	likes    int
}

type Comments struct {
	id       int
	postID   int
	authorID int
  	text string
	created  string
	likes    int
}

func main() {
	fmt.Println("1")
	newDB := createDatabase()
	addUser(newDB, 0, "buter", "bat@mail.ru", "123", "1212")
	rows, _ := newDB.Query("SELECT id,username, email, password, created FROM users")
	var id int
	var username, email, password, created string
	for rows.Next() {
		rows.Scan(&id, &username, &email, &password, &created)
		fmt.Printf("%d: %s %s %s %s\n", id, username, email, password, created)
	}
	// deleteUser(newDB, 0)
}

func createDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./newDB.db")
	checkErr(err)

	// if there is no such table, it will be created with the following properties
	statementForUsers, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users ( id INTEGER PRIMARY KEY, username TEXT NOT NULL UNIQUE, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, created TEXT )")
	statementForUsers.Exec()
	statementForPosts, _ := db.Prepare("CREATE TABLE IF NOT EXISTS posts ( id INTEGER PRIMARY KEY, authorID INTEGER PRIMARY KEY, created TEXT, likes INTEGER )")
	statementForPosts.Exec()
	// statementForComments, _ := db.Prepare("CREATE TABLE IF NOT EXISTS comments ( id INTEGER PRIMARY KEY,postID INTEGER, authorID INTEGER, created TEXT, likes INTEGER )")
	// statementForComments.Exec()
	return db
}

func addUser(db *sql.DB, id int, username string, email string, password string, created string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users (id, username, email, password, created) VALUES (?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	_, err := stmt.Exec(id, username, email, password, created)
	checkErr(err)
	tx.Commit()
}
func addPost(db *sql.DB, id int, authorID int, created string, likes int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO posts (id, authorID, created, likes) VALUES (?,?,?,?)")

	_, err := stmt.Exec(id, authorID, created, likes)
	checkErr(err)
	tx.Commit()
}
func deleteUser(db *sql.DB, id int) {
	sid := strconv.Itoa(id) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM users WHERE id=?")
	_, err := stmt.Exec(sid)
	checkErr(err)
	tx.Commit()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

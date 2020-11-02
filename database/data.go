package main

import (
	"database/sql"
	"fmt"
	"strconv"
	// "reflect"
	// "../vars"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)
func main() {
	// fmt.Println("1")
	newDB := createDatabase()
	addUser(newDB, createdUID(), "buterbrot", "bat@mail.ru", encryptPassword("abc"), "123456")
	fmt.Println(newDB)

}

func createDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./newDB.db")
	checkErr(err)
	createUser(db)
	createPost(db)
	createComments(db)
	return db
}
func createUser(db *sql.DB) {
	// db, _ := sql.Open("sqlite3", "./newDB.db")
	statementForUsers, err := db.Prepare(` 
	
	CREATE TABLE IF NOT EXISTS "users" (
		"id" UID PRIMARY KEY, 
		"username" TEXT NOT NULL UNIQUE, 
		"email" TEXT NOT NULL UNIQUE, 
		"password" BLOB, 
		"created" TEXT 
		); 
	
	`)
	checkErr(err)
	statementForUsers.Exec()
	
}
func createPost(db *sql.DB) {
	
	statementForPosts, err :=db.Prepare(` 
	
	CREATE TABLE IF NOT EXISTS "posts" ( 
		"id" UID PRIMARY KEY, 
		"authorID" UID , 
		"title" TEXT, 
		"created" TEXT, 
		"category" TEXT,
		"likes" INTEGER 
		);
		COMMIT;
	`)
	checkErr(err)
	statementForPosts.Exec()
}
func createComments(db *sql.DB)  {
	statementForComments, err :=db.Prepare(` 
	
	CREATE TABLE IF NOT EXISTS "comments" ( 
		"id" UID PRIMARY KEY, 
		"postID" UID ,
		"authorID" UID,
		"text" TEXT, 
		"created" TEXT, 
		"likes" INTEGER 
		);
		COMMIT;
	`)
	checkErr(err)
	statementForComments.Exec()
}
func addUser(db *sql.DB,id uuid.UUID, username string, email string, password []byte, created string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users (id, username, email, password, created) VALUES (?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	_, err := stmt.Exec(id,username, email, password, created)
	checkErr(err)
	tx.Commit()
}
func addPost(db *sql.DB, id uuid.UUID, authorID uuid.UUID, title string, created string, category string, likes int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO posts (id, authorID, title, created, category, likes) VALUES (?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	_, err := stmt.Exec(id, authorID, title, created, category, likes)
	checkErr(err)
	tx.Commit()
}
func addComment(db *sql.DB,id uuid.UUID, postID uuid.UUID, authorID uuid.UUID, text string, created string, likes int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO comments (id, postID, authorID, text, created, likes) VALUES (?,?,?,?,?)")
	_, err := stmt.Exec(id, postID, authorID, text, created, likes)
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
func createdUID() uuid.UUID {
	u1 := uuid.Must(uuid.NewV4())
	return u1
}
func encryptPassword(pas string) []byte {
	enc, err := bcrypt.GenerateFromPassword([]byte(pas), bcrypt.MinCost) // def is 4
	checkErr(err)
	return enc
}
func checkPassword(enc []byte, pas string) bool {
	return bcrypt.CompareHashAndPassword(enc, []byte(pas)) == nil
}

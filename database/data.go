package data

import (
	"database/sql"
	"fmt"

	// "reflect"
	"../vars"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var AllUsers []vars.User
var AllPost []vars.Post

//CreateDatabase creates db
func CreateDatabase() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "./mainDB.db")
	db.Exec("PRAGMA foreign_keys = ON")
	CheckErr(err)
	CreateUsers(db)
	CreatePosts(db)
	CreateComments(db)
	CreateSessions(db)
	fmt.Println("DATABASE CREATED")

	return db
}

// CreateUsers s
func CreateUsers(db *sql.DB) {
	// db, _ := sql.Open("sqlite3", "./newDB.db")
	statementForUsers, err := db.Prepare(` 
	
	CREATE TABLE IF NOT EXISTS "users" (
		"id" UID NOT NULL PRIMARY KEY, 
		"username" TEXT NOT NULL UNIQUE,
		"email" TEXT NOT NULL UNIQUE, 
		"password" BLOB, 
		"created" TEXT 
		); 
	
	`)
	CheckErr(err)
	statementForUsers.Exec()

}
func CreatePosts(db *sql.DB) {

	statementForPosts, err := db.Prepare(` 
	
	CREATE TABLE IF NOT EXISTS "posts" ( 
		"id" UID NOT NULL PRIMARY KEY, 
		"authorID" UID,
		"title" TEXT, 
		"text" TEXT, 
		"created" TEXT, 
		"category" TEXT,
		"likes" INTEGER,
		FOREIGN KEY(authorID)REFERENCES users(id)
		);
		
	`)
	CheckErr(err)
	statementForPosts.Exec()
}
func CreateComments(db *sql.DB) {
	statementForComments, err := db.Prepare(` 
	
	CREATE TABLE IF NOT EXISTS "comments" ( 
		"id" UID NOT NULL PRIMARY KEY, 
		"postID" UID ,
		"authorID" UID ,
		"text" TEXT, 
		"created" TEXT, 
		"likes" INTEGER,
		FOREIGN KEY(postID)REFERENCES posts(id),
		FOREIGN KEY(authorID)REFERENCES users(id)
		);
		
	`)
	CheckErr(err)
	statementForComments.Exec()
}
func CreateSessions(db *sql.DB) {

	statementForPosts, err := db.Prepare(` 
	
	CREATE TABLE IF NOT EXISTS "session" ( 
		"id" UID NOT NULL PRIMARY KEY, 
		"userID" UID,
		"cookieID" UID
		);
		
	`)
	CheckErr(err)
	statementForPosts.Exec()
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
func CreatedUID() uuid.UUID {
	u1 := uuid.Must(uuid.NewV4())
	return u1
}
func CheckPassword(enc []byte, pas string) bool {
	return bcrypt.CompareHashAndPassword(enc, []byte(pas)) == nil
}
func DbConn() (db *sql.DB) {

	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		panic(err.Error())
	}
	return db
}

package data

import (
	"database/sql"
	"fmt"

	// "reflect"
	// "../vars"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

//CreateDatabase creates db
func CreateDatabase() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "../mainDB.db")

	CheckErr(err)
	CreateUser(db)
	CreatePost(db)
	CreateComments(db)
	fmt.Println("1")
	return db
}
func CreateUser(db *sql.DB) {
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
func CreatePost(db *sql.DB) {

	statementForPosts, err := db.Prepare(` 
	
	CREATE TABLE IF NOT EXISTS "posts" ( 
		"id" UID NOT NULL PRIMARY KEY, 
		"authorID" UID, 
		"title" TEXT, 
		"created" TEXT, 
		"category" TEXT,
		"likes" INTEGER 
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
		"likes" INTEGER 
		);
		
	`)
	CheckErr(err)
	statementForComments.Exec()
}

func AddPost(db *sql.DB, id uuid.UUID, authorID uuid.UUID, title string, created string, category string, likes int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO posts (id, authorID, title, created, category, likes) VALUES (?,?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	_, err := stmt.Exec(id, authorID, title, created, category, likes)
	CheckErr(err)
	tx.Commit()
}
func AddComment(db *sql.DB, id uuid.UUID, postID uuid.UUID, authorID uuid.UUID, text string, created string, likes int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO comments (id, postID, authorID, text, created, likes) VALUES (?,?,?,?,?)")
	_, err := stmt.Exec(id, postID, authorID, text, created, likes)
	CheckErr(err)
	tx.Commit()
}

//DeleteUser deletes

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

package database

import (
	"database/sql"
	"strconv"

	// "reflect"
	// "../vars"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// fmt.Println("1")
	newDB := CreateDatabase()
	AddUser(newDB, CreatedUID(), "buterbrot", "bat@mail.ru", EncryptPassword("abc"), "123456")

}

//CreateDatabase creates db
func CreateDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./newDB.db")
	checkErr(err)
	CreateUser(db)
	CreatePost(db)
	CreateComments(db)
	return db
}
func CreateUser(db *sql.DB) {
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
func CreatePost(db *sql.DB) {

	statementForPosts, err := db.Prepare(` 
	
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
func CreateComments(db *sql.DB) {
	statementForComments, err := db.Prepare(` 
	
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
func AddUser(db *sql.DB, id uuid.UUID, username string, email string, password []byte, created string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users (id, username, email, password, created) VALUES (?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	_, err := stmt.Exec(id, username, email, password, created)
	checkErr(err)
	tx.Commit()
}
func AddPost(db *sql.DB, id uuid.UUID, authorID uuid.UUID, title string, created string, category string, likes int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO posts (id, authorID, title, created, category, likes) VALUES (?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	_, err := stmt.Exec(id, authorID, title, created, category, likes)
	checkErr(err)
	tx.Commit()
}
func AddComment(db *sql.DB, id uuid.UUID, postID uuid.UUID, authorID uuid.UUID, text string, created string, likes int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO comments (id, postID, authorID, text, created, likes) VALUES (?,?,?,?,?)")
	_, err := stmt.Exec(id, postID, authorID, text, created, likes)
	checkErr(err)
	tx.Commit()
}

//DeleteUser deletes
func DeleteUser(db *sql.DB, id int) {
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
func CreatedUID() uuid.UUID {
	u1 := uuid.Must(uuid.NewV4())
	return u1
}
func EncryptPassword(pas string) []byte {
	enc, err := bcrypt.GenerateFromPassword([]byte(pas), bcrypt.MinCost) // def is 4
	checkErr(err)
	return enc
}
func CheckPassword(enc []byte, pas string) bool {
	return bcrypt.CompareHashAndPassword(enc, []byte(pas)) == nil
}

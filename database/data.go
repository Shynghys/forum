package database

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	// _ "github.com/satori/go.uuid"
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
	name     string
	created  string
	comments Comments
	likes    int
}

type Comments struct {
	id       int
	postID   int
	authorID int
	text     string
	created  string
	likes    int
}

// func main() {
// 	fmt.Println("1")
// 	newDB := createDatabase()
// 	addUser(newDB, 0, "buter", "bat@mail.ru", "123", "1212")
// 	rows, _ := newDB.Query("SELECT id,username, email, password, created FROM users")
// 	var id int
// 	var username, email, password, created string
// 	for rows.Next() {
// 		rows.Scan(&id, &username, &email, &password, &created)
// 		fmt.Printf("%d: %s %s %s %s\n", id, username, email, password, created)
// 	}
// 	// deleteUser(newDB, 0)
// }

//CreateDatabase creates db
func CreateDatabase() *sql.DB {
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

//AddUser adds user
func AddUser(db *sql.DB, id int, username string, email string, password string, created string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users (id, username, email, password, created) VALUES (?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	_, err := stmt.Exec(id, username, email, password, created)
	checkErr(err)
	tx.Commit()
}

//AddPost adds user
func AddPost(db *sql.DB, id int, authorID int, created string, likes int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO posts (id, authorID, created, likes) VALUES (?,?,?,?)")

	_, err := stmt.Exec(id, authorID, created, likes)
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

// func d0ssans_part() {
// 	// if there is no such file, it will be created
// 	db, _ := sql.Open("sqlite3", "./mydb.db")

// 	// if there is no such table, it will be created with the following properties
// 	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users ( id INTEGER PRIMARY KEY, nickname TEXT NOT NULL UNIQUE, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL )")
// 	statement.Exec()

// 	// to insert variables
// 	statement, _ = db.Prepare("INSERT INTO users (nickname, email, password) VALUES (?, ?, ?)")
// 	enc, err := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.MinCost) // def is 4
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	statement.Exec("J", "s@gmai.cm", string(enc)) // exact these values will not work cause nickname "Jane" is already exists

// 	// bcryption, NIL is Right Pswrd
// 	fmt.Println(bcrypt.CompareHashAndPassword(enc, []byte("123123")))

// 	// to parse variables
// 	rows, _ := db.Query("SELECT id, nickname, email, password FROM users")
// 	var id int
// 	var nickname, email, password string
// 	for rows.Next() {
// 		rows.Scan(&id, &nickname, &email, &password)
// 		fmt.Printf("%d: %s %s %s\n", id, nickname, email, password)
// 	}
// }

package data

import (
	"database/sql"

	"../vars"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// create User
func AddUser(db *sql.DB, username string, email string, password string, created string) {
	tx, _ := db.Begin()
	passwordEnc := EncryptPassword(password)
	id := CreatedUID()
	stmt, _ := tx.Prepare("INSERT INTO users (id, username, email, password, created) VALUES (?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	_, err := stmt.Exec(id, username, email, passwordEnc, created)
	CheckErr(err)
	tx.Commit()
}

// get User from table. PS. ----> add child table for posts and comments
func GetUsers(db *sql.DB, id2 uuid.UUID) vars.User {
	rows, err := db.Query("select * from users")
	CheckErr(err)
	for rows.Next() {
		var tempUser vars.User
		err =
			rows.Scan(&tempUser.ID, &tempUser.Username, &tempUser.Email, &tempUser.Password, &tempUser.Created /*, &tempUser.posts, &tempUser.comments*/)
		CheckErr(err)
		if tempUser.ID == id2 {
			return tempUser
		}
	}
	return vars.User{}
}

// func updateUser(db *sql.DB, id2 int, username string, surname string, age int, university string) {
// 	sage := strconv.Itoa(age) // int to string
// 	sid := strconv.Itoa(id2)  // int to string
// 	tx, _ := db.Begin()
// 	stmt, _ := tx.Prepare("update testTable set username=?,surname=?,age=?,university=? where id=?")
// 	_, err := stmt.Exec(username, surname, sage, university, sid)
// 	checkError(err)
// 	tx.Commit()
// }

// delete User ss
func DeleteUser(db *sql.DB, id uuid.UUID) {

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM users WHERE id=?")
	_, err := stmt.Exec(id)
	CheckErr(err)
	tx.Commit()
}
func EncryptPassword(pas string) []byte {
	enc, err := bcrypt.GenerateFromPassword([]byte(pas), bcrypt.MinCost) // def is 4
	CheckErr(err)
	return enc
}

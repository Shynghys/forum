package data

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"github.com/shynghys/forum/vars"
	"golang.org/x/crypto/bcrypt"
)

// create User
func CreateUser(user *vars.User) {
	db := DbConn()
	defer db.Close()

	// var newUser vars.User
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	passwordEnc := string(EncryptPassword(user.Password))
	// id := CreatedUID()
	user.ID = CreatedUID()
	// newUser.ID = id
	// newUser.Username = user.Username
	// newUser.Email = user.Email
	// newUser.Created = user.Created
	user.Created = time.Now().Format(time.RFC1123)
	// AllUsers = append(AllUsers, newUser)
	stmt, _ := tx.Prepare("INSERT INTO users (id, username, email, password, created) VALUES (?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	stmt.Exec(user.ID, user.Username, user.Email, passwordEnc, user.Created)
	// if err != nil {
	// 	fmt.Println("This user already exist")
	// }
	tx.Commit()
}
func ReadUser(id2 uuid.UUID) vars.User {
	db := DbConn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users")
	CheckErr(err)
	defer rows.Close()
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
func GetUsername(id2 uuid.UUID) string {
	db := DbConn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users")
	CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		var tempUser vars.User
		err =
			rows.Scan(&tempUser.ID, &tempUser.Username, &tempUser.Email, &tempUser.Password, &tempUser.Created /*, &tempUser.posts, &tempUser.comments*/)
		CheckErr(err)
		if tempUser.ID == id2 {
			return tempUser.Username
		}
	}
	return "anonymos"
}
func UpdateUser(toChange vars.User) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update users set username=?, email=?, password=? where id=?")
	_, err := stmt.Exec(toChange.Username, toChange.Email, toChange.Password, toChange.ID)
	CheckErr(err)
	tx.Commit()
}
func DeleteUser(id uuid.UUID) {
	db := DbConn()
	defer db.Close()
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

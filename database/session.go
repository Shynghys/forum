package data

import (
	"log"

	"../vars"
	uuid "github.com/satori/go.uuid"
)

func CreateSession(session vars.Session) {

	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()

	_, err := db.Exec("INSERT INTO session (sessionID, userID) VALUES (?,?)", session.SessionID, session.UserID)
	// stmt, err := tx.Prepare("INSERT INTO posts (id, authorID, title, text, created, category, likes) VALUES (?,?,?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	// _, err := stmt.Exec(post.ID, post.AuthorID, post.Title, post.Text, post.Created, post.Category, post.Likes)
	// fmt.Println(result)
	if err != nil {
		log.Println(err)
	}

	// CheckErr(err)

	tx.Commit()
}
func DeleteSession(id uuid.UUID) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("DELETE FROM session WHERE sessionID=?")
	_, err := stmt.Exec(id)
	CheckErr(err)
	tx.Commit()
}

func DeleteSessionByID(id uuid.UUID) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("DELETE FROM session WHERE userID=?")
	_, err := stmt.Exec(id)
	CheckErr(err)
	tx.Commit()
}

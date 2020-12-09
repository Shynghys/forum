package data

import (
	"fmt"
	"log"

	"../vars"
	uuid "github.com/satori/go.uuid"
)

func CreateSession(session vars.Session) {

	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()

	id := CreatedUID()

	result, err := db.Exec("INSERT INTO session (id, userID, cookieID) VALUES (?,?,?)", id, session.UserID, session.SessionID)
	// stmt, err := tx.Prepare("INSERT INTO posts (id, authorID, title, text, created, category, likes) VALUES (?,?,?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	// _, err := stmt.Exec(post.ID, post.AuthorID, post.Title, post.Text, post.Created, post.Category, post.Likes)
	fmt.Println(result)
	if err != nil {
		log.Println(err)
	}

	// CheckErr(err)
	fmt.Println("Session created!!!!1")
	tx.Commit()
}
func DeleteSession(id uuid.UUID) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM session WHERE id=?")
	_, err := stmt.Exec(id)
	CheckErr(err)
	tx.Commit()
}

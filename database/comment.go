package data

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shynghys/forum/vars"
)

func CreateComment(comment vars.Comment) uuid.UUID {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	id := CreatedUID()
	comment.Created = time.Now().Format(time.RFC1123)
	// comment.AuthorID = ""

	stmt, _ := tx.Prepare("INSERT INTO comments (id, postID, authorID, author, text, created, likes, dislikes) VALUES (?,?,?,?,?,?,?,?)")
	_, err := stmt.Exec(id, comment.PostID, comment.AuthorID, comment.Author, comment.Text, comment.Created, comment.Likes, comment.Dislikes)

	CheckErr(err)
	tx.Commit()
	return id
}

func ReadComment(id2 uuid.UUID) vars.Comment {
	db := DbConn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM comments")
	CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		var tempComment vars.Comment
		err =
			rows.Scan(&tempComment.ID, &tempComment.PostID, &tempComment.AuthorID, &tempComment.Text, &tempComment.Created, &tempComment.Likes, tempComment.Dislikes)
		CheckErr(err)
		if tempComment.ID == id2 {
			return tempComment
		}
	}
	return vars.Comment{}
}

func UpdateComment(toChange vars.Comment) {
	db := DbConn()
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("UPDATE comments SET text=?, created=?, likes=? WHERE id=?")
	_, err := stmt.Exec(toChange.Text, toChange.Created, toChange.Likes, toChange.ID)
	CheckErr(err)
	tx.Commit()
}
func DeleteComment(id uuid.UUID) {

	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM comments WHERE id=?")
	_, err := stmt.Exec(id)
	CheckErr(err)
	tx.Commit()
}

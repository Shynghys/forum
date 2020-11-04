package data

import (
	"database/sql"

	"../vars"
	uuid "github.com/satori/go.uuid"
)

func AddComment(db *sql.DB, comment vars.Comment) {
	tx, _ := db.Begin()
	id := CreatedUID()
	stmt, _ := tx.Prepare("INSERT INTO comments (id, postID, authorID, text, created, likes) VALUES (?,?,?,?,?,?)")
	_, err := stmt.Exec(id, comment.PostID, comment.AuthorID, comment.Text, comment.Created, comment.Likes)
	CheckErr(err)
	tx.Commit()
}
func GetComment(db *sql.DB, id2 uuid.UUID) vars.Comment {
	rows, err := db.Query("SELECT * FROM comments")
	CheckErr(err)
	for rows.Next() {
		var tempComment vars.Comment
		err =
			rows.Scan(&tempComment.ID, &tempComment.PostID, &tempComment.AuthorID, &tempComment.Text, &tempComment.Created, &tempComment.Likes)
		CheckErr(err)
		if tempComment.ID == id2 {
			return tempComment
		}
	}
	return vars.Comment{}
}
func UpdateComment(db *sql.DB, toChange vars.Comment) {

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("UPDATE comments SET text=?, created=?, likes=? WHERE id=?")
	_, err := stmt.Exec(toChange.Text, toChange.Created, toChange.Likes, toChange.ID)
	CheckErr(err)
	tx.Commit()
}
func DeleteComment(db *sql.DB, id uuid.UUID) {

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM comments WHERE id=?")
	_, err := stmt.Exec(id)
	CheckErr(err)
	tx.Commit()
}

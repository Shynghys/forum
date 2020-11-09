package data

import (
	"database/sql"

	"../vars"
	uuid "github.com/satori/go.uuid"
)

func CreatePost(db *sql.DB, post vars.Post) {
	tx, _ := db.Begin()
	id := CreatedUID()
	stmt, _ := tx.Prepare("INSERT INTO posts (id, authorID, title, created, category, likes) VALUES (?,?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	_, err := stmt.Exec(id, post.AuthorID, post.Title, post.Created, post.Category, post.Likes)
	CheckErr(err)
	tx.Commit()
}
func ReadPost(db *sql.DB, id2 uuid.UUID) vars.Post {
	rows, err := db.Query("SELECT * FROM posts")
	CheckErr(err)
	for rows.Next() {
		var tempPost vars.Post
		err =
			rows.Scan(&tempPost.ID, &tempPost.AuthorID, &tempPost.Title, &tempPost.Created, &tempPost.Category, &tempPost.Likes /*, &tempPost.posts, &tempPost.comments*/)
		CheckErr(err)
		if tempPost.ID == id2 {
			return tempPost
		}
	}
	return vars.Post{}
}
func UpdatePost(db *sql.DB, toChange vars.Post) {

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("UPDATE posts SET title=?, category=?, likes=? WHERE id=?")
	_, err := stmt.Exec(toChange.Title, toChange.Category, toChange.Likes, toChange.ID)
	CheckErr(err)
	tx.Commit()
}
func DeletePost(db *sql.DB, id uuid.UUID) {

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM posts WHERE id=?")
	_, err := stmt.Exec(id)
	CheckErr(err)
	tx.Commit()
}

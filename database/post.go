package data

import (
	"fmt"
	"log"

	"../vars"
	uuid "github.com/satori/go.uuid"
)

func ReadAllPosts() []vars.Post {

	db := DbConn()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts")
	CheckErr(err)

	posts := []vars.Post{}

	for rows.Next() {
		var tempPost vars.Post
		err =
			rows.Scan(&tempPost.ID, &tempPost.AuthorID, &tempPost.Title, &tempPost.Text, &tempPost.Created, &tempPost.Category, &tempPost.Likes /*, &tempPost.posts, &tempPost.comments*/)
		CheckErr(err)
		posts = append(posts, tempPost)
	}
	fmt.Println(posts)
	return posts
}
func CreatePost(post *vars.Post) {

	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()

	id := CreatedUID()

	result, err := db.Exec("INSERT INTO posts (id, authorID, title, text, created, category, likes) VALUES (?,?,?,?,?,?,?)", id, post.AuthorID, post.Title, post.Text, post.Created, post.Category, post.Likes)
	// stmt, err := tx.Prepare("INSERT INTO posts (id, authorID, title, text, created, category, likes) VALUES (?,?,?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	// _, err := stmt.Exec(post.ID, post.AuthorID, post.Title, post.Text, post.Created, post.Category, post.Likes)
	fmt.Println(result)
	if err != nil {
		log.Println(err)
	}

	// CheckErr(err)
	fmt.Println("Post created!!!!1")
	tx.Commit()
}
func ReadPost(title string) vars.Post {
	db := DbConn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM posts")
	CheckErr(err)
	for rows.Next() {
		var tempPost vars.Post
		err =
			rows.Scan(&tempPost.ID, &tempPost.AuthorID, &tempPost.Title, &tempPost.Text, &tempPost.Created, &tempPost.Category, &tempPost.Likes /*, &tempPost.posts, &tempPost.comments*/)
		CheckErr(err)
		if tempPost.Title == title {
			return tempPost
		}
	}
	return vars.Post{}
}
func UpdatePost(title string, toChange vars.Post) vars.Post {
	db := DbConn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM posts")
	CheckErr(err)
	for rows.Next() {
		var tempPost vars.Post
		err =
			rows.Scan(&tempPost.ID, &tempPost.AuthorID, &tempPost.Title, &tempPost.Text, &tempPost.Created, &tempPost.Category, &tempPost.Likes /*, &tempPost.posts, &tempPost.comments*/)
		CheckErr(err)
		if tempPost.Title == title {
			return toChange
		}
	}
	return vars.Post{}
}
func DeletePost(id uuid.UUID) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM posts WHERE id=?")
	_, err := stmt.Exec(id)
	CheckErr(err)
	tx.Commit()
}

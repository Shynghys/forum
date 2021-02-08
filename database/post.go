package data

import (
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shynghys/forum/vars"
)

func ReadAllPosts() []vars.Post {

	db := DbConn()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts")
	CheckErr(err)
	defer rows.Close()

	posts := []vars.Post{}

	for rows.Next() {
		var tempPost vars.Post
		err =
			rows.Scan(&tempPost.ID, &tempPost.AuthorID, &tempPost.Author, &tempPost.Title, &tempPost.Image, &tempPost.Text, &tempPost.Created, &tempPost.Category, &tempPost.Likes, &tempPost.Dislikes /*, &tempPost.posts, &tempPost.comments*/)
		CheckErr(err)
		posts = append(posts, tempPost)
	}

	return posts
}

func CreatePost(post *vars.Post) uuid.UUID {
	db := DbConn()
	defer db.Close()

	tx, _ := db.Begin()
	id := CreatedUID()
	post.Created = time.Now().Format(time.RFC1123)
	_, err := db.Exec("INSERT INTO posts (id, authorID,author, title, image, text, created, category, likes, dislikes) VALUES (?,?,?,?,?,?,?,?,?,?)", id, post.AuthorID, post.Author, post.Title, post.Image, post.Text, post.Created, post.Category, post.Likes, post.Dislikes)
	// stmt, err := tx.Prepare("INSERT INTO posts (id, authorID, title, text, created, category, likes) VALUES (?,?,?,?,?,?,?)")
	// stmt.Exec(id, username, email, password, created)
	// _, err := stmt.Exec(post.ID, post.AuthorID, post.Title, post.Text, post.Created, post.Category, post.Likes)
	if err != nil {
		log.Fatalln(err)
	}

	tx.Commit()
	return id
}
func ReadPost(title string) vars.Post {
	db := DbConn()
	defer db.Close()

	rows1, err := db.Query("SELECT * FROM comments")
	CheckErr(err)
	var Comms []vars.Comment
	for rows1.Next() {
		var Comm vars.Comment
		// var Created sql.NullInt64

		err =
			rows1.Scan(&Comm.ID, &Comm.PostID, &Comm.AuthorID, &Comm.Author, &Comm.Text, &Comm.Created, &Comm.Likes, &Comm.Dislikes /*, &Comm.Like , &tempPost.posts, &tempPost.comments*/)
		CheckErr(err)
		needID, _ := uuid.FromString(title)
		if Comm.PostID == needID {
			// &tempPost.Comments = comments
			Comms = append(Comms, Comm)

		}
	}
	rows, err := db.Query("SELECT * FROM posts")
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var tempPost vars.Post
		err =
			rows.Scan(&tempPost.ID, &tempPost.AuthorID, &tempPost.Author, &tempPost.Title, &tempPost.Image, &tempPost.Text, &tempPost.Created, &tempPost.Category, &tempPost.Likes, &tempPost.Dislikes /*, &tempPost.posts, &tempPost.comments*/)
		CheckErr(err)
		needID, _ := uuid.FromString(title)
		if tempPost.ID == needID {
			tempPost.Comments = Comms
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
	defer rows.Close()

	for rows.Next() {
		var tempPost vars.Post
		err =
			rows.Scan(&tempPost.ID, &tempPost.AuthorID, &tempPost.Author, &tempPost.Title, &tempPost.Image, &tempPost.Text, &tempPost.Created, &tempPost.Category, &tempPost.Likes, &tempPost.Dislikes /*, &tempPost.posts, &tempPost.comments*/)
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

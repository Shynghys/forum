package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	uuid "github.com/satori/go.uuid"
	database "github.com/shynghys/forum/database"
	db "github.com/shynghys/forum/database"
	"github.com/shynghys/forum/vars"
)

// PostsHandler gets posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {

	// if r.Method == "GET" {
	tmpl := template.Must(template.ParseFiles("templates/homepage.html"))
	AllPosts := db.ReadAllPosts()

	tmpl.Execute(w, AllPosts)

	// }

}

// CreatePost gets post by id
func CreatePost(w http.ResponseWriter, r *http.Request) {

	if !(r.URL.Path == "/posts/create") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		c, _ := r.Cookie(COOKIE_NAME)

		if c != nil {
			needCookie := GetUserByCookie(r)
			if needCookie == "" {
				cookieID, err := uuid.FromString(GetCookie(r, COOKIE_NAME))
				if err != nil {
					// fmt.Printf("Something went wrong: %s", err)
					return
				}
				database.DeleteSession(cookieID)
				DeleteCookie(w, r)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				tmpl := template.Must(template.ParseFiles("templates/createpost.html"))
				tmpl.Execute(w, nil)
			}
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else if r.Method == "POST" {
		fmt.Println(r.FormValue("title"))
		c, _ := r.Cookie(COOKIE_NAME)
		if c != nil {
			needCookie := GetUserByCookie(r)
			if needCookie == "" {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				if err := r.ParseForm(); err != nil {
					fmt.Fprintf(w, "ParseForm() err: %v", err)
					return
				}
				tmpl := template.Must(template.ParseFiles("templates/createpost.html"))
				var ErrorMsg Message
				r.Body = http.MaxBytesReader(w, r.Body, vars.MAX_UPLOAD_SIZE)
				file, fileHeader, err := r.FormFile("file") // extracting file from the request body
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				// reading the file

				if fileHeader.Size > vars.MAX_UPLOAD_SIZE { // parsing request body into form data

					ErrorMsg.Msg = "The uploaded file is too big. Please choose an file that's less than 20MB in size"
					tmpl.Execute(w, ErrorMsg)

					return // if the file is too big

				}

				defer file.Close()
				buff := make([]byte, 512)
				_, err = file.Read(buff) // reading the file into the buffer
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				filetype := http.DetectContentType(buff) // detecting file type
				if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/gif" && filetype != "image/svg+xml" {
					ErrorMsg.Msg = "The provided file format is not allowed. Please upload a JPEG, GIF or PNG image"
					tmpl.Execute(w, ErrorMsg)

				}
				_, err = file.Seek(0, io.SeekStart) // makes sure the file is read from the start
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				err = os.MkdirAll("./uploads", os.ModePerm) // makes uploads directory if it doesnt exist
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				// Create a new file in the uploads directory
				newName := time.Now().UnixNano()
				dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", newName, filepath.Ext(fileHeader.Filename))) // creates and names the file for copying
				fmt.Println(newName)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				defer dst.Close()

				_, err = io.Copy(dst, file) // copies the file content into the new one in the uploads folder
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				categories := r.FormValue("movies") + " " + r.FormValue("books") + " " + r.FormValue("games")
				details := vars.Post{
					Title: r.FormValue("title"),
					Text:  r.FormValue("text"),
					Image: fmt.Sprintf("%d%s", newName, filepath.Ext(fileHeader.Filename)),

					Category: categories,
				}

				if !isReadable(details.Text) || !isReadable(details.Title) {
					http.Redirect(w, r, "/posts/create", http.StatusSeeOther)
					return
				}

				details.AuthorID, _ = uuid.FromString(GetUserByCookie(r))
				details.Author = db.GetUsername(details.AuthorID)
				id := db.CreatePost(&details)
				db.CreateLike(id)
				db.CreateDislike(id)

				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

}

// ReadPost gets post by id
func ReadPost(w http.ResponseWriter, r *http.Request) {

	_, err := os.Open("./mainDB.db")
	if err != nil {
		ErrorHandler(w, r, 500)
		return
	}

	type page struct {
		UserDetails PageDetails
		Posts       vars.Post
	}
	var b page
	c, _ := r.Cookie(COOKIE_NAME)
	if c != nil {
		b.UserDetails.UserIn = true
		needCookie, _ := uuid.FromString(GetUserByCookie(r))
		findUser := db.ReadUser(needCookie)
		b.UserDetails.UserName = findUser.Username
		b.UserDetails.UserID = findUser.ID
	}

	title := r.URL.Query().Get("id")

	b.Posts = db.ReadPost(title)

	if b.Posts.Title == "" {
		ErrorHandler(w, r, 400)
		return
	}

	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/show-post.html"))

		tmpl.Execute(w, b)
	} else if r.Method == "POST" {

		// tmpl := template.Must(template.ParseFiles("templates/show-post.html"))
		// if r.Method != http.MethodPost {
		// 	tmpl.Execute(w, nil)
		// 	return
		// }
		details := vars.Comment{
			AuthorID: b.UserDetails.UserID,
			Author:   b.UserDetails.UserName,
			PostID:   b.Posts.ID,
			Text:     r.FormValue("text"),
		}

		like := r.FormValue("like")
		dislike := r.FormValue("dislike")

		if like != "" || dislike != "" {

			details.Author = db.GetUsername(details.AuthorID)

			if like != "" {
				likeUUID, _ := uuid.FromString(like)
				db.LikeBtn(likeUUID, details.AuthorID)
			} else if dislike != "" {
				dislikeUUID, _ := uuid.FromString(dislike)
				db.DislikeBtn(dislikeUUID, details.AuthorID)
			}

		} else if isReadable(details.Text) {
			id := db.CreateComment(details)
			db.CreateLike(id)
			db.CreateDislike(id)
		}

		path := r.URL.Path + "?id=" + title

		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, path, http.StatusSeeOther)

	}
}

// UpdatePost gets post by id
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// r.Body.Read()
	title := r.URL.Query().Get("title")
	// if !(r.URL.Path == "/posts/update") {
	// 	ErrorHandler(w, r, http.StatusNotFound)
	// 	return
	// }
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/edit-post.html"))
		Post := db.ReadPost(title)

		tmpl.Execute(w, Post)

	}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		details := vars.Post{
			Title:    r.FormValue("title"),
			Text:     r.FormValue("text"),
			Category: r.FormValue("category"),
		}

		db.UpdatePost(title, details)

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

// DeletePost gets post by id
func DeletePost(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/posts/{id}/delete") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

		tmpl.Execute(w, nil)

	}

	if r.Method == "POST" {
		tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		// db.DeletePost(id)

		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

func isReadable(str string) bool {
	for _, v := range str {
		if v > ' ' {
			return true
		}
	}
	return false
}

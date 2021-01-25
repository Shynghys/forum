package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	db "../database/"
	"../vars"
	uuid "github.com/satori/go.uuid"
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
		tmpl := template.Must(template.ParseFiles("templates/createpost.html"))
		tmpl.Execute(w, nil)
	}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		var a []string

		if r.FormValue("movies/serials") != "" {
			a = append(a, r.FormValue("movies/serials"))
		}
		if r.FormValue("books") != "" {
			a = append(a, r.FormValue("books"))
		}
		if r.FormValue("games") != "" {
			a = append(a, r.FormValue("games"))
		}
		categories := strings.Join(a, ",")

		details := vars.Post{
			Title:    r.FormValue("title"),
			Text:     r.FormValue("text"),
			Category: categories,
		}
		// fmt.Println(details)
		details.AuthorID, _ = uuid.FromString(GetUserByCookie(r))
		db.CreatePost(&details)

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

// ReadPost gets post by id
func ReadPost(w http.ResponseWriter, r *http.Request) {
	type page struct {
		UserDetails PageDetails
		Posts       vars.Post
	}
	var b page
	var IsUserin PageDetails
	IsUserin.UserIn = false
	c, _ := r.Cookie(COOKIE_NAME)
	if c != nil {
		IsUserin.UserIn = true
		needCookie, _ := uuid.FromString(GetUserByCookie(r))
		findUser := db.ReadUser(needCookie)
		IsUserin.UserName = findUser.Username

	}
	b.UserDetails = IsUserin

	title := r.URL.Query().Get("id")
	fmt.Println("===================")
	if r.Method == "GET" {
		fmt.Println("dasdad")
		// if !(r.URL.Path == "/posts/{id}") {
		// 	ErrorHandler(w, r, http.StatusNotFound)
		// 	return
		// }

		fmt.Println(title)
		tmpl := template.Must(template.ParseFiles("templates/show-post.html"))
		Post := db.ReadPost(title)
		b.Posts = Post
		fmt.Println(b)
		fmt.Println("readpost done")
		tmpl.Execute(w, b)

	}
	if r.Method == "POST" {
		fmt.Println("CREATING COMMENT")
		tmpl := template.Must(template.ParseFiles("templates/show-post.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		postID, _ := uuid.FromString(r.URL.Query().Get("id"))
		details := vars.Comment{
			PostID: postID,
			Text:   r.FormValue("text"),
		}
		details.AuthorID, _ = uuid.FromString(GetUserByCookie(r))
		fmt.Println(details)
		db.CreateComment(details)

		// do something with details
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)

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

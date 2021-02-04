package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	db "github.com/shynghys/forum/database"
	"github.com/shynghys/forum/vars"
)

// CommentsHandler gets Comments
func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/comments/") {
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
		details := vars.Comment{
			Text: r.FormValue("text"),
		}
		// db.AddUser(details)

		// do something with details
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

// CreateComment gets post by id
func CreateComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww")
	// if !(r.URL.Path == "/comments/create/") {
	// 	ErrorHandler(w, r, http.StatusNotFound)
	// 	return
	// }

	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))
		tmpl.Execute(w, nil)
	}

	if r.Method == "POST" {
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

// ReadComment gets Comment by id
func ReadComment(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/comments/{id}") {
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
		details := vars.Comment{
			Text: r.FormValue("text"),
		}
		// db.GetComment(details)

		// do something with details
		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

// UpdateComment gets Comment by id
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/comments/{id}/update") {
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
		details := vars.Comment{
			Text: r.FormValue("text"),
		}
		// db.UpdateComment(details)

		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

// DeleteComment gets Comment by id
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/comments/{id}/delete") {
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
		details := vars.Comment{
			Text: r.FormValue("text"),
		}
		// db.AddUser(details)

		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

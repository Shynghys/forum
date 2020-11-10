package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	db "../database/"
	"../vars"
)

// PostsHandler gets posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {

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
		details := vars.User{
			Email:    r.FormValue("email"),
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		// db.GetPosts(details)

		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

// CreatePost gets post by id
func CreatePost(w http.ResponseWriter, r *http.Request) {

	// if !(r.URL.Path == "/posts/create") {
	// 	ErrorHandler(w, r, http.StatusNotFound)
	// 	return
	// }
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/createpost.html"))

		tmpl.Execute(w, nil)

	}

	if r.Method == "POST" {
		fmt.Println("wwwwwwwwwoooooo")
		// tmpl := template.Must(template.ParseFiles("templates/createpost.html"))
		// if r.Method != http.MethodPost {
		// 	tmpl.Execute(w, nil)
		// 	return
		// }
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Println("000000000000")
		details := &vars.Post{

			Title:    r.FormValue("title"),
			Text:     r.FormValue("text"),
			Category: r.FormValue("category"),
			Created:  "0",
			Likes:    12,
		}
		fmt.Println(details)
		fmt.Println("111111111111111")
		db.CreatePost(details)
		fmt.Println("2222222222222")
		// tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

// ReadPost gets post by id
func ReadPost(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/posts/{id}") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))
		//Post:=GetPost(id)
		tmpl.Execute(w, nil)

	}

}

// UpdatePost gets post by id
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/posts/{id}/update") {
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
		details := vars.Post{
			Title:    r.FormValue("title"),
			Text:     r.FormValue("text"),
			Category: r.FormValue("category"),
		}
		// db.UpdatePost(details)

		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
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

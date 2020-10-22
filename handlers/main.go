package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"../vars"

	"github.com/gorilla/mux"
)

// Handler does smth
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// SignInHandler signs in
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := vars.User{
		Email: r.FormValue("email"),
		// Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// do something with details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})

}

// SignUpHandler signs up
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := vars.User{
		Email:    r.FormValue("email"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	password := r.FormValue("confirmation-password")
	if details.Password != password {
		fmt.Println("did not match")
	}
	// do something with details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})

}

// UserHandler gets user by id
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

// PostsHandler gets post by id
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

// PostHandler gets post by id
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

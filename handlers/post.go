package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// PostsHandler gets post by id
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

// PostHandler gets post by id
// func PostHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
// }

// CreatePost gets post by id
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

// ReadPost gets post by id
func ReadPost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

// UpdatePost gets post by id
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

// DeletePost gets post by id
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

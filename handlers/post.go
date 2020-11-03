package handlers

import (
	"html/template"
	"net/http"
)

// PostsHandler gets post by id
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)
	// fmt.Fprintf(w, "You've requested the user: id = %s \n")
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
	// db.AddPost()
	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)

}

// ReadPost gets post by id
func ReadPost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)
}

// UpdatePost gets post by id
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)
}

// DeletePost gets post by id
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)
}

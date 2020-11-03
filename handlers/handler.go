package handlers

import (
	"html/template"
	"net/http"
)

var templates *template.Template
var utilPattern string

// LoadTemplates func
func LoadTemplates(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))
	utilPattern = pattern
}

// ExecuteTemplate func
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	LoadTemplates(utilPattern)
	templates.ExecuteTemplate(w, tmpl, data)
}

// NewRouter s
func NewRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", Handler)
	r.HandleFunc("/sign_in", SignInHandler)
	r.HandleFunc("/sign_up", SignUpHandler)
	// r.HandleFunc("/user/{id}", UserHandler)

	r.HandleFunc("/users", UsersHandler)
	r.HandleFunc("/users/{id}/create", CreateUser)
	r.HandleFunc("/users/{id}", ReadUser)
	r.HandleFunc("/users/{id}/update", UpdateUser)
	r.HandleFunc("/users/{id}/delete", DeleteUser)

	r.HandleFunc("/posts", PostsHandler)
	r.HandleFunc("/posts/{id}/create", CreatePost)
	r.HandleFunc("/posts/{id}", ReadPost)
	r.HandleFunc("/posts/{id}/update", UpdatePost)
	r.HandleFunc("/posts/{id}/delete", DeletePost)

	r.HandleFunc("/comments", CommentsHandler)
	r.HandleFunc("/comments/{id}/create", CreateComment)
	r.HandleFunc("/comments/{id}", ReadComment)
	r.HandleFunc("/comments/{id}/update", UpdateComment)
	r.HandleFunc("/comments/{id}/delete", DeleteComment)

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	return r
}

func checkErr(err error) error {
	return err
}

// Handler does smth
func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("templates/main.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	tmpl.Execute(w, struct{ Success bool }{true})
	// http.Redirect(w, r, "/", 200)
}

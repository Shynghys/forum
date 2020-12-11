package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	db "../database/"
	"../vars"
	uuid "github.com/satori/go.uuid"
)

var templates *template.Template
var utilPattern string

// Temps is for handling error tempaltes
var Temps *template.Template

// Error is for tempaltes
var Error vars.ErrorStruct

type PageDetails struct {
	UserIn   bool
	UserName string
	AllPosts []vars.Post
}

// NewRouter s
func NewRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", Handler)
	r.HandleFunc("/sign-in", SignInHandler)
	r.HandleFunc("/sign-up", SignUpHandler)
	r.HandleFunc("/logout", LogoutHandler)
	// r.HandleFunc("/user/{id}", UserHandler)

	r.HandleFunc("/users", UsersHandler)
	r.HandleFunc("/users/create", CreateUser)
	r.HandleFunc("/users/{id}", ReadUser)
	r.HandleFunc("/users/{id}/update", UpdateUser)
	r.HandleFunc("/users/{id}/delete", DeleteUser)

	// r.HandleFunc("/posts", PostsHandler)
	r.HandleFunc("/posts/create", CreatePost)
	r.HandleFunc("/posts", ReadPost)
	r.HandleFunc("/posts/edit", UpdatePost)
	r.HandleFunc("/posts/delete", DeletePost)

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
	if err != nil {
		return err
	}
	return nil
}

// Handler does smth
func Handler(w http.ResponseWriter, r *http.Request) {
	var isUserin PageDetails
	isUserin.UserIn = false
	c, _ := r.Cookie(COOKIE_NAME)
	// if err != nil {
	// 	panic(err)
	// }
	tmpl := template.Must(template.ParseFiles("templates/homepage.html"))
	if c != nil {
		isUserin.UserIn = true
		needCookie, _ := uuid.FromString(GetUserByCookie(r))
		findUser := db.ReadUser(needCookie)
		isUserin.UserName = findUser.Username

	}
	fmt.Println(isUserin)
	isUserin.AllPosts = db.ReadAllPosts()

	tmpl.Execute(w, isUserin)

	// http.Redirect(w, r, "/", 200)
}

//ErrorHandler handles error
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	Temps := template.Must(template.ParseGlob("templates/error.html"))

	if status == 404 {
		Error.Status = 404
		Error.StatusDefinition = "Not found"

	} else if status == 500 {
		Error.Status = 500
		Error.StatusDefinition = "Internal server problem"
	} else if status == 400 {
		Error.Status = 400
		Error.StatusDefinition = "Bad request"
	}

	Temps.ExecuteTemplate(w, "error.html", Error)
}

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

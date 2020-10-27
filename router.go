package main

import (
	"net/http"

	"./handlers"
)

//MakeRouter handles routing
func MakeRouter(w http.ResponseWriter, r *http.Request) {
	// http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/sign_in", handlers.SignInHandler)
	http.HandleFunc("/sign_up", handlers.SignUpHandler)

	http.HandleFunc("/user/{id}", handlers.UserHandler)

	http.HandleFunc("/posts", handlers.PostsHandler)
	http.HandleFunc("/posts/{title}", handlers.CreatePost)
	// http.HandleFunc("/posts/{title}", handlers.ReadPost)
	// http.HandleFunc("/posts/{title}", handlers.UpdatePost)
	// http.HandleFunc("/posts/{title}", handlers.DeletePost)

	// http.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	// http.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
}

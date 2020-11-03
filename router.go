package main

import (
	"net/http"
)

//MakeRouter handles routing
func MakeRouter(w http.ResponseWriter, r *http.Request) {
	// http.HandleFunc("/", h.Handler)
	// http.HandleFunc("/sign_in", h.SignInHandler)
	// http.HandleFunc("/sign_up", h.SignUpHandler)

	// http.HandleFunc("/user/{id}", h.UserHandler)

	// http.HandleFunc("/posts", h.PostsHandler)
	http.HandleFunc("/posts/{id}/create", h.CreatePost)
	http.HandleFunc("/posts/{id}", h.ReadPost)
	http.HandleFunc("/posts/{id}/update", h.UpdatePost)
	http.HandleFunc("/posts/{id}/delete", h.DeletePost)

	// http.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	// http.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
}

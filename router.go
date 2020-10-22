package main

import (
	"./handlers"
	"github.com/gorilla/mux"
)

func makeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Handler)
	r.HandleFunc("/sign_in", handlers.SignInHandler)
	r.HandleFunc("/sign_up", handlers.SignUpHandler)
	r.HandleFunc("/", handlers.Handler)
	r.HandleFunc("/user/{id}", handlers.UserHandler)
	r.HandleFunc("/posts", handlers.PostsHandler)
	r.HandleFunc("/post/{id}", handlers.PostHandler)
	// r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	return r
}

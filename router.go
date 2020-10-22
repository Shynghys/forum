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

	r.HandleFunc("/user/{id}", handlers.UserHandler)

	r.HandleFunc("/posts", handlers.PostsHandler)
	r.HandleFunc("/posts/{title}", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{title}", handlers.ReadPost).Methods("GET")
	r.HandleFunc("/posts/{title}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{title}", handlers.DeletePost).Methods("DELETE")

	// r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	return r
}

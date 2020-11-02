package main

import (
	"fmt"
	"log"
	"net/http"

	h "./handlers"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {

	// Starting web server

	// http.HandleFunc("/", MakeRouter)
	http.HandleFunc("/", h.Handler)
	http.HandleFunc("/sign_in", h.SignInHandler)
	http.HandleFunc("/sign_up", h.SignUpHandler)

	http.HandleFunc("/user/{id}", h.UserHandler)

	http.HandleFunc("/posts", h.PostsHandler)
	http.HandleFunc("/posts/{title}", h.CreatePost)
	fmt.Println("serving..............")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

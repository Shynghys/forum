package main

import (
	"fmt"
	"log"
	"net/http"

	h "./handlers"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	//Create db
	// newDb := db.CreateDatabase()
	// fmt.Println(newDb)

	// Starting web server

	// http.HandleFunc("/", MakeRouter)
	h.LoadTemplates("templates/*.html")
	r := h.NewRouter()
	http.Handle("/", r)
	fmt.Println("serving..............")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

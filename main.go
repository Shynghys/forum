package main

import (
	"fmt"
	"log"
	"net/http"

	db "./database"

	h "./handlers"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	//Create db
	db.CreateDatabase()
	// fmt.Println(newDb)

	// Starting web server

	// http.HandleFunc("/", MakeRouter)
	h.LoadTemplates("templates/tmpl/*.html")
	r := h.NewRouter()
	http.Handle("/", r)
	fmt.Println("serving..............")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

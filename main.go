package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/shynghys/forum/database"

	h "github.com/shynghys/forum/handlers"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	//Create db
	db.CreateDatabase()

	h.LoadTemplates("templates/tmpl/*.html")
	r := h.NewRouter()
	http.Handle("/", r)
	fmt.Println("serving..............")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

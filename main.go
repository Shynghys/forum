package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	db "github.com/shynghys/forum/database"

	h "github.com/shynghys/forum/handlers"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	//Create db
	db.CreateDatabase()

	h.LoadTemplates("templates/*.html")
	r := h.NewRouter()
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("serving..............")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

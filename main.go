package main

import (
	"fmt"
	"log"
	"net/http"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {

	// Starting web server
	http.HandleFunc("/", MakeRouter)

	fmt.Println("serving..............")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

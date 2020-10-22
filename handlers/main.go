package handlers

import (
	"fmt"
	"net/http"
)

// Handler does smth
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

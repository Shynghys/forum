package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// UserHandler gets user by id
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

// // userHandlers gets user by id
// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	db, _ := sql.Open("sqlite3", "./m.db")
// 	userID := r.Header.Get("X-HashText-User-ID")

// 	row := db.QueryRow(`SELECT name, credit FROM "user" WHERE user_id = $1`, userID)

// 	var name string
// 	var credit int
// 	err := row.Scan(&name, &credit)
// 	switch {
// 	case err == sql.ErrNoRows:
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	case err != nil:
// 		log.Printf("Query to look up user failed: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	sendJSONResponse(w, userDocument{UserID: userID, Name: name, Credit: credit})
// }

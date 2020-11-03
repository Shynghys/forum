package handlers

import (
	"fmt"
	"html/template"
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

// UsersHandler gets Users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)
	// fmt.Fprintf(w, "You've requested the user: id = %s \n")
}

// CreateUser gets User by id
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// db.AddUser()
	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)

}

// ReadUser gets User by id
func ReadUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)
}

// UpdateUser gets User by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)
}

// DeleteUser gets User by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)
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

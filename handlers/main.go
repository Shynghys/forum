package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"../vars"

	"github.com/gorilla/mux"
)

// Handler does smth
func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}

// SignInHandler signs in
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/sign-in.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := vars.User{
		Email: r.FormValue("email"),
		// Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// do something with details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})

}

// SignUpHandler signs up
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := vars.User{
		Email:    r.FormValue("email"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	password := r.FormValue("confirmation-password")
	if details.Password != password {
		fmt.Println("did not match")
	}
	// do something with details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})

}

// UserHandler gets user by id
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "You've requested the user: id = %s \n", id)
}

// // userHandler gets user by id
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

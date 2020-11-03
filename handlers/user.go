package handlers

import (
	"html/template"
	"net/http"
<<<<<<< HEAD

	"../vars"
)

=======
)

// UserHandler gets user by id
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

>>>>>>> refs/remotes/origin/master
// UsersHandler gets Users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

	tmpl.Execute(w, nil)
	// fmt.Fprintf(w, "You've requested the user: id = %s \n")
}

// CreateUser gets User by id
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

		tmpl.Execute(w, nil)

	}

	if r.Method == "POST" {
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
		// db.AddUser(details)
		// password := r.FormValue("confirmation-password")
		// if details.Password != password {
		// 	fmt.Println("did not match")
		// }

		// do something with details
		_ = details
		fmt.Println(details.Email)
		// do something with details
		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//   saveChoice(r.Form["choices"])
		//   http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}
}

// ReadUser gets User by id
func ReadUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

		tmpl.Execute(w, nil)

	}

	if r.Method == "POST" {
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
		// db.AddUser(details)
		// password := r.FormValue("confirmation-password")
		// if details.Password != password {
		// 	fmt.Println("did not match")
		// }

		// do something with details
		_ = details
		fmt.Println(details.Email)
		// do something with details
		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//   saveChoice(r.Form["choices"])
		//   http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}
}

// UpdateUser gets User by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

		tmpl.Execute(w, nil)

	}

	if r.Method == "POST" {
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
		// db.AddUser(details)
		// password := r.FormValue("confirmation-password")
		// if details.Password != password {
		// 	fmt.Println("did not match")
		// }

		// do something with details
		_ = details
		fmt.Println(details.Email)
		// do something with details
		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//   saveChoice(r.Form["choices"])
		//   http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}
}

// DeleteUser gets User by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))

		tmpl.Execute(w, nil)

	}

	if r.Method == "POST" {
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
		// db.AddUser(details)
		// password := r.FormValue("confirmation-password")
		// if details.Password != password {
		// 	fmt.Println("did not match")
		// }

		// do something with details
		_ = details
		fmt.Println(details.Email)
		// do something with details
		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//   saveChoice(r.Form["choices"])
		//   http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}
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

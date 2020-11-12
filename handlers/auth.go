package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	data "../database"
	"../vars"
)

// SignInHandler signs in
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/sign-in") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		// tmpl, err := template.New("base").ParseFiles("templates/tmpl/sign-in.html", "templates/tmpl/base.html")
		tmpl := template.Must(template.ParseFiles("templates/sign-in.html"))

		// check your err

		// if err != nil {
		// panic(err)
		// }
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		tmpl.Execute(w, nil)
	}

	if r.Method == "POST" {
		tmpl := template.Must(template.ParseFiles("templates/sign-in.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		login := r.FormValue("login") // can be username or email
		// Username: r.FormValue("username"),
		password := r.FormValue("password")

		uuid := checkAll(login, password)
		if uuid == "" {
			log.Fatal("Username or password is incorrect.")
		}

		fmt.Println(uuid)
		// do something with details

		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//   saveChoice(r.Form["choices"])
		//   http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}
}

// SignUpHandler signs up
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/sign-up") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

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

		var isEmailUsed, isUsernameUsed bool
		isEmailUsed = checkEmail(details.Email) != ""
		isUsernameUsed = checkUsername(details.Username) != ""

		if isEmailUsed && isUsernameUsed {
			fmt.Println("these email and username are already in use.")
		} else if isEmailUsed {
			fmt.Println("This email is already in use.")
		} else if isUsernameUsed {
			fmt.Println("This username is already in use.")
		} else {
			// db := data.CreateDatabase()
			data.CreateUser(details)
			fmt.Println("You are cool.")
		}

		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//   saveChoice(r.Form["choices"])
		//   http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}
}

func checkEmail(email string) string {
	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT id FROM users WHERE email LIKE ?", email)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	var id string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id)
		log.Println("UUID is: ", id)
	}

	return id
}

func checkUsername(username string) string {
	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT id FROM users WHERE username LIKE ?", username)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	var id string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id)
		log.Println("UUID is: ", id)
	}

	return id
}

func checkPassword(password string) string {
	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT id FROM users WHERE email LIKE ?", password)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	var id string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id)
		log.Println("UUID is: ", id)
	}

	return id
}

func checkAll(data, password string) string {
	idEmail := checkEmail(data)
	idUsername := checkUsername(data)
	idPassword := checkPassword(password)

	if (idEmail == idPassword) || (idPassword == idUsername) {
		if idEmail == "" {
			return idUsername
		}
		return idEmail
	}
	return ""
}

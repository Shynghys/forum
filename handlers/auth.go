package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	data "../database"
	"../vars"
	"golang.org/x/crypto/bcrypt"
)

// SignInHandler signs in
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/sign-in") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

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

		// Username: r.FormValue("username"),
		type details struct {
			Login    string // can be username or email
			Password string
		}

		data := details{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
		}

		uuid := checkAll(db, data.Login, data.Password)
		if uuid == "" {
			http.Redirect(w, r, "/sign-up", http.StatusSeeOther) // something was wrong
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther) // need find idea how to send uuid...
		}
		// do something with details

		// tmpl.Execute(w, struct{ Success bool }{true})
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

	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
		isEmailUsed = checkEmail(db, details.Email) != ""
		isUsernameUsed = checkUsername(db, details.Username) != ""

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

func checkEmail(db *sql.DB, email string) string {

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

func checkUsername(db *sql.DB, username string) string {

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

func checkAll(db *sql.DB, data, password string) string {
	idEmail := checkEmail(db, data)
	idUsername := checkUsername(db, data)

	// to figure out was it printed an e-mail or a username
	var uuid string
	if idEmail != "" {
		uuid = idEmail
	} else {
		uuid = idUsername
	}

	//nested func is to compare a printed code with a enc code in db
	isPasswordRight := func(uuid, pas string) bool {
		row, err := db.Query("SELECT id FROM users WHERE id LIKE ?", uuid)
		if err != nil {
			log.Fatal(err)
		}

		defer row.Close()

		var password string
		for row.Next() { // Iterate and fetch the records from result cursor
			row.Scan(&password)
			log.Println("UUID is: ", password)
		}

		return bcrypt.CompareHashAndPassword([]byte(password), []byte(pas)) == nil
	}

	if !isPasswordRight(uuid, password) {
		return ""
	}

	return uuid
}

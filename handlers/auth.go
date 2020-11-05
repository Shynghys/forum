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
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/sign-in.html"))

		tmpl.Execute(w, nil)

	}

	if r.Method == "POST" {
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

		fmt.Println(details.Email)
		// do something with details
		_ = details

		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//   saveChoice(r.Form["choices"])
		//   http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}
}

// SignUpHandler signs up
func SignUpHandler(w http.ResponseWriter, r *http.Request) {

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

		isEmailUsed := checkEmail(details.Email)
		isUsernameUsed := checkUsername(details.Username)

		if isEmailUsed && isUsernameUsed {
			fmt.Println("these email and username are already in use.")
		} else if isEmailUsed {
			fmt.Println("This email is already in use.")
		} else if isUsernameUsed {
			fmt.Println("This username is already in use.")
		} else {
			db := data.CreateDatabase()
			data.AddUser(db, details)
			fmt.Println("You are cool.")
		}

		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//   saveChoice(r.Form["choices"])
		//   http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}
}

func checkEmail(email string) bool {
	db, err := sql.Open("sqlite3", "./newDB.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT uuid FROM student WHERE email LIKE ?", email)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	var uuid string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&uuid)
		log.Println("Program is: ", uuid)
	}

	return uuid != ""
}

func checkUsername(username string) bool {
	db, err := sql.Open("sqlite3", "./newDB.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT uuid FROM student WHERE username LIKE ?", username)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	var uuid string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&uuid)
		log.Println("Program is: ", uuid)
	}

	return uuid != ""
}

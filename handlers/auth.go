package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

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

	db, err := sql.Open("sqlite3", "./newDB.db")
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

		row, err := db.Query("SELECT uuid FROM student WHERE email LIKE ?", details.Email)
		if err != nil {
			log.Fatal(err)
		}
		defer row.Close()
		for row.Next() { // Iterate and fetch the records from result cursor
			var program string
			row.Scan(&program)
			log.Println("Program is: ", program)
		}

		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//   saveChoice(r.Form["choices"])
		//   http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}
}

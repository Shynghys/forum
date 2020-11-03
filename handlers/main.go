package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"../vars"
)

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", Handler)
	r.HandleFunc("/sign_in", SignInHandler)
	r.HandleFunc("/sign_up", SignUpHandler)
	r.HandleFunc("/user/{id}", UserHandler)
	r.HandleFunc("/posts", PostsHandler)
	r.HandleFunc("/posts/{title}", CreatePost)

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	return r
}

func checkErr(err error) error {
	return err
}

// Handler does smth
func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("templates/main.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}

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

package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	database "github.com/shynghys/forum/database"
	"github.com/shynghys/forum/vars"
	"golang.org/x/crypto/bcrypt"
)

type details struct {
	Login    string // can be username or email
	Password string
}

type Message struct {
	Msg string
}

const COOKIE_NAME = "my_cookie"

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/sign-in") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		ErrorHandler(w, r, 500)
		return
	}
	defer db.Close()

	tmpl := template.Must(template.ParseFiles("templates/sign-in.html"))
	var msg Message

	if r.Method == "GET" {
		tmpl.Execute(w, msg)
	} else if r.Method == "POST" {

		data := details{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
		}
		if data.Login == "" && data.Password == "" {
			msg.Msg = "2"
			tmpl.Execute(w, msg)
		}
		getUUID := checkAll(db, data.Login, data.Password)

		if getUUID == "" {
			http.Redirect(w, r, "/sign-up", http.StatusSeeOther) // something was wrong

		} else if getUUID == "error500" {
			ErrorHandler(w, r, 500)
		} else if getUUID == "wrong password" {
			msg.Msg = "1"
			tmpl.Execute(w, msg)
		} else {
			userid, _ := uuid.FromString(getUUID)
			sessionid := database.CreatedUID()
			database.DeleteSessionByID(userid)
			newSession := vars.Session{
				UserID:    userid,
				SessionID: sessionid,
			}
			database.CreateSession(newSession)
			cookie := &http.Cookie{
				Name:    COOKIE_NAME,
				Value:   sessionid.String(),
				Expires: time.Now().Add(60 * time.Minute),
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusSeeOther) // need find idea how to send uuid...
		}

	}
}

// SignUpHandler signs up
//done
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/sign-up") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		ErrorHandler(w, r, 500)
		return
	}
	defer db.Close()
	tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))
	var msg Message

	if r.Method == "GET" {
		tmpl.Execute(w, msg)
	}

	if r.Method == "POST" {
		// t := time.Now()
		details := vars.User{
			Username: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
			// Created:  t.Format(time.RFC1123),
		}
		// fmt.Println(details.Password)
		var isEmailUsed, isUsernameUsed bool
		isEmailUsed = checkEmail(db, details.Email) != ""
		isUsernameUsed = checkUsername(db, details.Username) != ""
		if isEmailUsed && isUsernameUsed {
			msg.Msg = "0"
			// msg.Msg = "These username and email are already in use"
			tmpl.Execute(w, msg)
			return
		}
		if isEmailUsed {
			msg.Msg = "1"
			// msg.Msg = "This email is already in use"
			tmpl.Execute(w, msg)
			return
		}
		if isUsernameUsed {
			msg.Msg = "2"
			// msg.Msg = "This username is already in use"
			tmpl.Execute(w, msg)
			return
		}

		database.CreateUser(&details)

		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)

	}
}

//done
func checkEmail(db *sql.DB, email string) string {

	row, err := db.Query("SELECT id FROM users WHERE email LIKE ?", email)
	if err != nil {
		// log.Fatal(err)
		return "error500"
	}

	defer row.Close()

	var id string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id)
		log.Println("UUID is: ", id)
	}

	return id
}

//done
func checkUsername(db *sql.DB, username string) string {

	row, err := db.Query("SELECT id FROM users WHERE username LIKE ?", username)
	if err != nil {
		// log.Fatal(err)
		return "error500"
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
	if idEmail == "error500" {
		return "error500"
	}
	idUsername := checkUsername(db, data)
	if idUsername == "error500" {
		return "error500"
	}

	// to figure out was it printed an e-mail or a username
	var uuid string
	if idEmail != "" {
		uuid = idEmail
	} else if idUsername != "" {
		uuid = idUsername
	} else {
		return ""
	}

	// fmt.Println("----Checkall UUID---------------")
	// fmt.Println(uuid)

	//nested func is to compare a printed code with a enc code in db
	isPasswordRight := func(uuid, pas string) bool {
		row, err := db.Query("SELECT password FROM users WHERE id LIKE ?", uuid)
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
		return "wrong password"
	}

	return uuid
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/logout") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cookieID, err := uuid.FromString(GetCookie(r, COOKIE_NAME))
	if err != nil {
		// fmt.Printf("Something went wrong: %s", err)
		return
	}
	database.DeleteSession(cookieID)
	DeleteCookie(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	// fmt.Println(cookieID)

}

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type student struct {
	id       string
	fullname string
	code     string
	program  program
}

type program struct {
	bachelor string
	master   string
	phd      string
}

func main() {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database

	// freshman := student{fullname: "Tom Holland", code: "007", program: "PHD"}
	// insertStudent(sqliteDatabase, freshman.code, freshman.fullname, freshman.program)
	// DISPLAY INSERTED RECORDS
	// displayStudents(sqliteDatabase)

	http.HandleFunc("/", indexhandle)

	// filesystem := http.FileServer(http.Dir("static"))
	// http.Handle("/", filesystem)

	http.ListenAndServe(":8080", nil)
	str := " hello "
	fmt.Println(len(strings.Split(str, " ")))
}

//indexhandle is function to implement  a web server
func indexhandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	var freshmeat student
	switch r.Method {
	case "GET": // no any client/server communication
		t.ExecuteTemplate(w, "index.html", freshmeat)
	case "POST": // send
		//handle inputs
		freshmeat.code = r.FormValue("code")
		freshmeat.fullname = r.FormValue("fullName")
		var a []string

		if r.FormValue("bachelor") != "" {
			a = append(a, r.FormValue("bachelor"))
		}
		if r.FormValue("master") != "" {
			a = append(a, r.FormValue("master"))
		}
		if r.FormValue("PHD") != "" {
			a = append(a, r.FormValue("PHD"))
		}
		if a == nil {
			fmt.Println("EPMTY")
			break
		}
		shit := strings.Join(a, ",")

		fmt.Println(shit)

		fmt.Println(len(strings.Split(shit, ",")))
		t.Execute(w, freshmeat) // execute with data above
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

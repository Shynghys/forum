package handlers

import (
	"database/sql"
	"log"
	"net/http"
)

func GetCookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return c.Value
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   COOKIE_NAME,
		MaxAge: -1}
	http.SetCookie(w, &c)
}

func GetUserByCookie(r *http.Request) string {
	db, err := sql.Open("sqlite3", "./mainDB.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	c, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		return ""
	}
	row, err := db.Query("SELECT userID FROM session WHERE sessionID LIKE ?", c.Value)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var id string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id)
	}
	return id
}

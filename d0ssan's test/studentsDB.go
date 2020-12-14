package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertStudent(db *sql.DB, code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayStudents(db *sql.DB) {
	name := "Liana Kim"
	// name := "Alayna Armiage"
	row, err := db.Query("SELECT program FROM student WHERE name LIKE ?", name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(row)
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var program string
		row.Scan(&program)
		log.Println("Program is: ", program)
	}
}

package main

import (
	d ".."
)

func main() {
	// fmt.Println("1")
	newDB := d.CreateDatabase()
	d.AddPost(newDB, d.CreatedUID(), d.CreatedUID(), "x", "x", "x", 3)
	d.AddUser(newDB, "buterbrot", "bat@mail.ru", "abc", "123456")

}

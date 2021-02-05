package main

import (
	d ".."
	v "../../vars"
)

func main() {
	// fmt.Println("1")
	newDB := d.CreateDatabase()
	var a v.User
	a.Username = "ButerBrot"
	a.Email = "batowka359@gmail.com"
	a.Password = "abc"
	a.Created = "21-03-1997"
	d.AddUser(newDB, a)
	// d.AddUser(newDB, "ray-bang8", "kanatkyzy@gmail.ru", "abc", "01-02-1998")
	// d.AddUser(newDB, "Dossan", "Yerdos@gmail.ru", "abc", "18-04-1997")
	// d.AddPost(newDB, d.CreatedUID(), d.AllUsers[0].ID, "x", "x", "x", 3)
	// d.UpdateUser(newDB, d.AllUsers[0].ID, "x", "x", "as")
	// d.DeleteUser(newDB, d.AllUsers[1].ID)
	// fmt.Println(d.GetUser(newDB, d.AllUsers[0].ID))

	// fmt.Println(d.AllUsers)
	// d.AddComment(newDB, d.CreatedUID(), d.CreatedUID(), d.CreatedUID(), "xx", "a", 3)

}

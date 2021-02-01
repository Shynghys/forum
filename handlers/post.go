package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	db "../database/"
	"../vars"
	uuid "github.com/satori/go.uuid"
)

// PostsHandler gets posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {

	// if r.Method == "GET" {
	tmpl := template.Must(template.ParseFiles("templates/homepage.html"))
	AllPosts := db.ReadAllPosts()

	tmpl.Execute(w, AllPosts)

	// }

}

// CreatePost gets post by id
func CreatePost(w http.ResponseWriter, r *http.Request) {

	if !(r.URL.Path == "/posts/create") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/createpost.html"))
		tmpl.Execute(w, nil)
	}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// var a []string

		// if r.FormValue("movies/serials") != "" {
		// 	a = append(a, r.FormValue("movies/serials"))
		// }
		// if r.FormValue("books") != "" {
		// 	a = append(a, r.FormValue("books"))
		// }
		// if r.FormValue("games") != "" {
		// 	a = append(a, r.FormValue("games"))
		// }

		// a = append(a, r.FormValue("select"))
		// fmt.Println(len(a))
		// fmt.Println(r.FormValue("select"))
		// categories := r.Form["select"][0]

		// if r.Form["select"][0] != "" {
		// 	a = append(a, "Movies")
		// }
		// if r.Form["select"][1] != "" {
		// 	a = append(a, "Games")
		// }
		// if r.Form["select"][2] != "" {
		// 	a = append(a, "Books")
		// }
		// if r.Form["movies"] != nil {
		// 	fmt.Println("fgfgfgfgfg")
		// }
		// fmt.Println(a)
		// categories := strings.Join(a, ",")
		categories := r.FormValue("select")
		fmt.Println("---------------Category---------------")
		fmt.Println(categories)
		details := vars.Post{
			Title:    r.FormValue("title"),
			Text:     r.FormValue("text"),
			Category: categories,
		}
		// fmt.Println(details)
		details.AuthorID, _ = uuid.FromString(GetUserByCookie(r))
		details.Author = db.GetUsername(details.AuthorID)
		id := db.CreatePost(&details)
		db.CreateLike(id)
		db.CreateDislike(id)

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

// ReadPost gets post by id
func ReadPost(w http.ResponseWriter, r *http.Request) {
	type page struct {
		UserDetails PageDetails
		Posts       vars.Post
	}
	var b page
	c, _ := r.Cookie(COOKIE_NAME)
	if c != nil {
		b.UserDetails.UserIn = true
		needCookie, _ := uuid.FromString(GetUserByCookie(r))
		findUser := db.ReadUser(needCookie)
		b.UserDetails.UserName = findUser.Username
		b.UserDetails.UserID = findUser.ID
	}

	title := r.URL.Query().Get("id")

	b.Posts = db.ReadPost(title)

	fmt.Println("===================")
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/show-post.html"))
		fmt.Println(b)
		tmpl.Execute(w, b)
	} else if r.Method == "POST" {
		fmt.Println("CREATING COMMENT")
		// tmpl := template.Must(template.ParseFiles("templates/show-post.html"))
		// if r.Method != http.MethodPost {
		// 	tmpl.Execute(w, nil)
		// 	return
		// }
		details := vars.Comment{
			AuthorID: b.UserDetails.UserID,
			Author:   b.UserDetails.UserName,
			PostID:   b.Posts.ID,
			Text:     r.FormValue("text"),
		}

		// details.AuthorID, _ = uuid.FromString(GetUserByCookie(r))

		// notAuthorised, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
		// if details.AuthorID == notAuthorised {
		// 	var err vars.ErrorStruct
		// 	tmpl := template.Must(template.ParseFiles("templates/error/index.html"))
		// 	err.Status = 401
		// 	err.StatusDefinition = "Not authorised"

		// 	tmpl.Execute(w, err)
		// }

		details.Author = db.GetUsername(details.AuthorID)

		like := r.FormValue("like")

		if like != "" {
			likeUUID, _ := uuid.FromString(like)
			fmt.Println("likeUUD", likeUUID)
			db.LikeBtn(likeUUID, details.AuthorID)
		}

		dislike := r.FormValue("dislike")

		if dislike != "" {
			dislikeUUID, _ := uuid.FromString(dislike)
			db.DislikeBtn(dislikeUUID, details.AuthorID)
		}

		fmt.Println(details)
		if details.Text != "" {
			id := db.CreateComment(details)
			db.CreateLike(id)
			db.CreateDislike(id)
		}

		// do something with details
		// _ = details
		path := r.URL.Path + "?id=" + title

		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, path, http.StatusSeeOther)

	}
}

// UpdatePost gets post by id
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// r.Body.Read()
	title := r.URL.Query().Get("title")
	// if !(r.URL.Path == "/posts/update") {
	// 	ErrorHandler(w, r, http.StatusNotFound)
	// 	return
	// }
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/edit-post.html"))
		Post := db.ReadPost(title)

		tmpl.Execute(w, Post)

	}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		details := vars.Post{
			Title:    r.FormValue("title"),
			Text:     r.FormValue("text"),
			Category: r.FormValue("category"),
		}

		db.UpdatePost(title, details)

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

// DeletePost gets post by id
func DeletePost(w http.ResponseWriter, r *http.Request) {
	if !(r.URL.Path == "/posts/{id}/delete") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
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

		// db.DeletePost(id)

		tmpl.Execute(w, struct{ Success bool }{true})
		// tmpl.Execute(w, struct{ Success bool }{true})
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

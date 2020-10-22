package vars

// User ss
type User struct {
	id       int
	Username string
	Email    string
	Password string
	posts    Post
	comments Comment
}

// Post ss
type Post struct {
	id       int
	authorID int
	comments Comment
	like     int
}

// Comment ss
type Comment struct {
	id       int
	postID   int
	authorID int
	like     int
}

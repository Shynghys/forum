package vars

// User ss
type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Created  string
	posts    Post
	comments Comment
}

type Post struct {
	ID       int
	AuthorID int
	Name     string
	Created  string
	comments Comment
	Likes    int
}

type Comment struct {
	ID       int
	PostID   int
	AuthorID int
	Message    string
	Created  string
	Likes    int
}
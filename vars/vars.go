package vars

import (
	uuid "github.com/satori/go.uuid"
)

// User ss
type User struct {
	ID       uuid.UUID
	Username string
	Email    string
	Password string
	Created  string
	posts    []Post
	comments []Comment
}

// Post ss
type Post struct {
	// id       string
	ID       uuid.UUID
	AuthorID uuid.UUID
	Title    string
	Category string
	Text     string
	Created  string
	Likes    int
	comments []Comment
}

// Comment ss
type Comment struct {
	ID       uuid.UUID
	PostID   uuid.UUID
	AuthorID uuid.UUID
	Text     string
	Created  string
	Likes    int
}

// ErrorStruct s
type ErrorStruct struct {
	Status           int
	StatusDefinition string
}

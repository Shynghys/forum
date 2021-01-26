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
	Dislikes int
	Comments []Comment
}

type Like struct {
}

// Comment ss
type Comment struct {
	ID       uuid.UUID
	PostID   uuid.UUID
	AuthorID uuid.UUID
	Text     string
	Created  string
	Like     []Like
	Likes    int
	Dislikes int
}

// ErrorStruct s
type ErrorStruct struct {
	Status           int
	StatusDefinition string
}

type Session struct {
	UserID    uuid.UUID
	SessionID uuid.UUID
}

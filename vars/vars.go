package vars

import (
	uuid "github.com/satori/go.uuid"
)

//MAX_UPLOAD_SIZE ss
const MAX_UPLOAD_SIZE int64 = 20 * 1024 * 1024 // 20mb
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
	Author   string
	Title    string
	Category string
	Text     string
	Image    string
	Created  string
	Likes    int
	Dislikes int
	Comments []Comment
	Liked    bool
}

// Comment ss
type Comment struct {
	ID       uuid.UUID
	PostID   uuid.UUID
	AuthorID uuid.UUID
	Author   string
	Text     string
	Created  string
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

type Like struct {
	ID        uuid.UUID
	AuthorsID interface{}
	Str       string
}

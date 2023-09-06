package models

import "github.com/google/uuid"

type Book struct {
	BaseModel
	Name    string    `json:"name" form:"name"`
	Auther  string    `json:"auther" form:"auther"`
	Edition string    `json:"edition" form:"edition"`
	Image   string    `json:"image" form:"image"`
	File    string    `json:"file" form:"file"`
	UserID  uuid.UUID `json:"userID" form:"userID"`
	User    User      `json:"user" form:"user"`
}

type Like struct {
	Book Book `json:"book" form:"book"`
	User User `json:"user" form:"user"`
}

type Comment struct {
	Book Book `json:"book" form:"book"`
	User User `json:"user" form:"user"`
	Body string
}

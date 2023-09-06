package models

import "github.com/google/uuid"

type Book struct {
	BaseModel
	Name    string    `json:"name" form:"name" gorm:"unique"`
	Auther  string    `json:"auther" form:"auther"`
	Edition string    `json:"edition" form:"edition"`
	Image   string    `json:"image" form:"image"`
	File    string    `json:"file" form:"file"`
	UserID  uuid.UUID `json:"userID" form:"userID"`
	User    User      `json:"user" form:"user"`
	Likes   []Like    `json:"likes" form:"likes"`
}

type Like struct {
	BookID uuid.UUID `json:"book" form:"book" `
	UserID uuid.UUID `json:"user" form:"user"`
}

type Comment struct {
	BaseModel
	BookID uuid.UUID `json:"book" form:"book" `
	UserID uuid.UUID `json:"user" form:"user"`
	Body   string    `json:"text" form:"text"`
}

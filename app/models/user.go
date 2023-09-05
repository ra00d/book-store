package models

// ALl attributtes to user and authentication and authorization stuff

import (
	"time"
)

type User struct {
	BaseModel
	Name        string `json:"name" form:"name"`
	Email       string
	Password    string
	Salt        string
	Phone       string
	RoleID      int
	Roles       Role         `gorm:"foreignKey:RoleID"`
	Permissions []Permission `json:"permissions" gorm:"many2many:user_permissions"`
	LastLogin   time.Time
	LastLoginIp string
	Enabled     bool
	rmemeberMe  string
}
type Role struct {
	ID   int `gorm:"primaryKey,autoIncrement" `
	Name string
}
type Permission struct {
	ID   int `gorm:"primaryKey,autoIncrement" `
	Name string
}

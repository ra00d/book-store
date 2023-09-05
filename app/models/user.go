package models

// ALl attributtes to user and authentication and authorization stuff

import (
	"time"
)

type User struct {
	BaseModel
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email" gorm:"unique"`
	Password    string `json:"password"`
	Salt        string `json:"-"`
	Phone       string `json:"phone"`
	RoleID      int
	Role        Role         `json:"role"`
	Permissions []Permission `json:"permissions" gorm:"many2many:user_permissions"`
	LastLogin   time.Time
	LastLoginIp string
	Enabled     bool `json:"enabled" gorm:"default:true"`
	rmemeberMe  string
}
type Role struct {
	ID   int    `gorm:"primaryKey,autoIncrement" `
	Name string `json:"name" gorm:"unique"`
}
type Permission struct {
	ID   int    `gorm:"primaryKey,autoIncrement" `
	Name string `json:"name" gorm:"unique"`
}

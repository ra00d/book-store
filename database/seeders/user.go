package seeders

import (
	"log"

	"github.com/ra00d/book_store/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	var (
		role = models.Role{
			Name: "admin",
		}

		permissions = []models.Permission{
			{Name: "create_user"},
			{Name: "create_book"},
			{Name: "delete_user"},
			{Name: "delete_book"},
			{Name: "update_user"},
			{Name: "update_book"},
		}
		user = models.User{
			Name:     "Admin",
			Email:    "admin@info.com",
			Password: string(hash),
		}
	)
	db.Create(&role)
	db.Create(&permissions)
	db.Create(&user)
}

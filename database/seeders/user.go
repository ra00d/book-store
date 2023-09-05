package seeders

import (
	"github.com/ra00d/book_store/app/models"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	var (
		role = models.Role{
			Name: "admin",
		}
        permissions=models.Permission{
            ID
        }
		user = models.User{
			Name:  "Admin",
			Email: "admin@info.com",
		}
	)
}

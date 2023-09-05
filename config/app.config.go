package config

import (
	// "github.com/gofiber/storage/sqlite3"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
	"github.com/ra00d/book_store/database"
	"gorm.io/gorm"
)

type AppConfig struct {
	Db           *gorm.DB
	SessionStore *session.Store
}

func (conf *AppConfig) Init() {
	conf.Db = database.Connect()
	storage := sqlite3.New(sqlite3.Config{
		Database: "./sessions.db",
		Table:    "session_data",
	})
	conf.SessionStore = session.New(session.Config{
		CookieSecure:   true,
		CookieSameSite: "strict",
		CookieHTTPOnly: true,
		KeyLookup:      "cookie:app_session",
		Storage:        storage,

		// source:         "header",
	})
}

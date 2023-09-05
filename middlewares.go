package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var ConfigDefault = monitor.Config{
	Title:      monitor.ConfigDefault.Title,
	Refresh:    monitor.ConfigDefault.Refresh,
	FontURL:    monitor.ConfigDefault.FontURL,
	ChartJsURL: monitor.ConfigDefault.ChartJsURL,
	CustomHead: monitor.ConfigDefault.CustomHead,
	APIOnly:    false,
	Next:       nil,
}

func SetFiberMiddlewares(app *fiber.App) {
	app.Get("/metrics", monitor.New(ConfigDefault))
	app.Use(requestid.New())
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		// Output: file,
		// Done: func(c *fiber.Ctx, logString []byte) {
		// 	db := GetAppDB(c)
		// 	db.Exec("INSERT INTO logs (content)VALUES (?)", string(logString))
		// },
	}))
	// app.Use(pprof.New())
	app.Use(recover.New())
	app.Use(limiter.New(limiter.Config{
		Max:        30,
		Expiration: 1 * time.Hour,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Render("limit", fiber.Map{})
		},
		LimiterMiddleware: limiter.ConfigDefault.LimiterMiddleware,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))
	// Initialize default config
	// app.Use(etag.New())
}

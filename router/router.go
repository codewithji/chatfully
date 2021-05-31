package router

import (
	"github.com/codewithji/chatfully/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())
	api.Get("/", handler.Home)
}

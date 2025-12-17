package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user-api/internal/handler"
)

// setup routes
func Setup(app *fiber.App, userHandler *handler.UserHandler) {
	// user routes
	app.Post("/users", userHandler.Create)
	app.Get("/users", userHandler.List)
	app.Get("/users/:id", userHandler.GetByID)
	app.Put("/users/:id", userHandler.Update)
	app.Delete("/users/:id", userHandler.Delete)
}

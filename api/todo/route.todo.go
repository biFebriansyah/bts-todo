package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func TodoRoute(app *fiber.App, db *sqlx.DB) {
	genre := app.Group("/todo")

	repos := NewRepo(db)
	handler := NewHandler(repos)

	genre.Get("/", handler.GetCards)
	genre.Post("/", handler.AddCard)
	// genre.Delete("/:uuid", handler.Delete)
}

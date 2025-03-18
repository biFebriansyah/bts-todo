package todo

import (
	"github.com/biFebriansyah/bts-todoapp/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func TodoRoute(app *fiber.App, db *sqlx.DB) {
	card := app.Group("/card")
	todo := app.Group("/todo")

	repos := NewRepo(db)
	handler := NewHandler(repos)

	//! cards
	card.Get("/", middleware.AuthMiddleware, handler.GetCards)
	card.Post("/", middleware.AuthMiddleware, handler.AddCard)
	card.Delete("/:uid", middleware.AuthMiddleware, handler.DeleteCards)

	//! todolist
	todo.Get("/", middleware.AuthMiddleware, handler.GetTodo)
	todo.Post("/", middleware.AuthMiddleware, handler.AddTodo)

	// genre.Delete("/:uuid", handler.Delete)
}

package todo

import (
	"github.com/biFebriansyah/bts-todoapp/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func TodoRoute(app *fiber.App, db *sqlx.DB) {
	todo := app.Group("/checklist")

	repos := NewRepo(db)
	handler := NewHandler(repos)

	todo.Get("/:cardId/item", middleware.AuthMiddleware, handler.GetTodo)
	todo.Post("/:cardId/item", middleware.AuthMiddleware, handler.AddTodo)
	todo.Get("/:cardId/item/:itemId", middleware.AuthMiddleware, handler.GetTodoId)
	todo.Delete("/:cardId/item/:itemId", middleware.AuthMiddleware, handler.DeleteTodo)
	todo.Put("/:cardId/item/:itemId", middleware.AuthMiddleware, handler.UpdateStatus)

	todo.Get("/", middleware.AuthMiddleware, handler.GetCards)
	todo.Post("/", middleware.AuthMiddleware, handler.AddCard)
	todo.Delete("/:uid", middleware.AuthMiddleware, handler.DeleteCards)
}

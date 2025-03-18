package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UserRoute(app *fiber.App, db *sqlx.DB) {
	genre := app.Group("/auth")

	repos := NewRepo(db)
	handler := NewHandler(repos)

	genre.Post("/signin", handler.SignIn)
	genre.Post("/signup", handler.SignUp)
	// genre.Delete("/:uuid", handler.Delete)
}

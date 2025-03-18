package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UserRoute(app *fiber.App, db *sqlx.DB) {
	auth := app.Group("/auth")

	repos := NewRepo(db)
	handler := NewHandler(repos)

	auth.Post("/signin", handler.SignIn)
	auth.Post("/signup", handler.SignUp)
	// genre.Delete("/:uuid", handler.Delete)
}

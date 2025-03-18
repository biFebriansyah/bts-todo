package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UserRoute(app *fiber.App, db *sqlx.DB) {

	repos := NewRepo(db)
	handler := NewHandler(repos)

	app.Post("/login", handler.SignIn)
	app.Post("/register", handler.SignUp)
}

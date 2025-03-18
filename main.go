package main

import (
	"log"

	"github.com/biFebriansyah/bts-todoapp/api/todo"
	"github.com/biFebriansyah/bts-todoapp/api/users"
	"github.com/biFebriansyah/bts-todoapp/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database := utils.NewDatabase()

	users.UserRoute(app, database.DB)
	todo.TodoRoute(app, database.DB)

	if err := app.Listen(":8083"); err != nil {
		log.Fatalf("Could not listen on 80803: %v\n", err)
	}
}

package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func main() {
	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Listen(":3000")
}

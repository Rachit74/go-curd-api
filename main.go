package main

import (
	"go-curd-api/database"
	"go-curd-api/models"
	"go-curd-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// connect database
	database.Connect()

	// migrate schema
	database.DB.AutoMigrate(&models.Book{})

	// setup routes
	routes.BookRoutes(app)

	// start server
	app.Listen(":3000")
}

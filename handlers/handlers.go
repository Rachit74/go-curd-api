package handlers

import (
	"go-curd-api/database"
	"go-curd-api/models"

	"github.com/gofiber/fiber/v2"
)

// Get all books
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)
	return c.JSON(books)
}

// Get single book by ID
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

// Create a new book
func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&book)
	return c.JSON(book)
}

// Update a book
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Save(&book)
	return c.JSON(book)
}

// Delete a book
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	database.DB.Delete(&book)
	return c.JSON(fiber.Map{"message": "Book deleted"})
}

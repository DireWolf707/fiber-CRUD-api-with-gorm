package book

import (
	"github.com/fiber-api/database"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.Conn
	var books []database.Book
	db.Find(&books)
	return c.JSON(books)
}
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Conn
	book := new(database.Book)
	db.Find(&book, id)
	return c.JSON(book)
}
func NewBook(c *fiber.Ctx) error {
	db := database.Conn
	book := new(database.Book)
	err := c.BodyParser(book)
	if err != nil {
		return c.Status(503).JSON(err)
	}
	db.Create(&book)
	return c.JSON(book)
}
func DeleteBooks(c *fiber.Ctx) error {
	db := database.Conn
	id := c.Params("id")
	book := new(database.Book)
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("Not Found!")
	}
	db.Delete(&book)
	return c.SendString("Book Deleted!")
}
func UpdateBooks(c *fiber.Ctx) error {
	db := database.Conn
	id := c.Params("id")
	book := new(database.Book)
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("Not Found!")
	}
	err := c.BodyParser(book)
	if err != nil {
		return c.Status(503).JSON(err)
	}
	db.Save(&book)
	return c.SendString("Book Updated!")
}

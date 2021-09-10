package main

import (
	"fmt"

	"github.com/fiber-api/book"
	"github.com/fiber-api/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Hello, World!") })
	app.Get("/api/v1/books", book.GetBooks)
	app.Post("/api/v1/books", book.NewBook)
	app.Get("/api/v1/books/:id", book.GetBook)
	app.Delete("/api/v1/books/:id", book.DeleteBooks)
	app.Put("/api/v1/books/:id", book.UpdateBooks)
}

func initDatabase() {
	var err error
	database.Conn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("Database Connected")
	database.Conn.AutoMigrate(&database.Book{})
	fmt.Printf("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":8080")
}

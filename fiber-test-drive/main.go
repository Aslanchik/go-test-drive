package main

import (
	"fmt"
	"go-test-drive/fiber-test-drive/book"
	"go-test-drive/fiber-test-drive/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()

	initDB()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Put("/api/v1/book", book.UpdateBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDB() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database Connection Success!")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated!")
}

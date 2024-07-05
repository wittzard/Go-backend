package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Harry Potter", Author: "J.K.Rowling"})
	books = append(books, Book{ID: 2, Title: "Metamorphosis", Author: "Frank Kafka"})

	//path   //Handlers
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)

	app.Listen(":8080")
}

// Handlers
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id")) //get value from /id

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookID {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Book not found")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)  // ตัวแทนรับ request
	c.BodyParser(book) //แปลง
	return c.JSON(book)
}

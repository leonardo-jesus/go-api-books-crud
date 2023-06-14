package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const BASE_URL string = "localhost:3333"

type book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Quantity    int32  `json:"quantity"`
}

type bookName struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var books = []book{
	{
		ID:          1,
		Title:       "Clean Code",
		Description: "A Handbook of Agile Software Craftsmanship",
		Author:      "Robert C. Martin",
		Quantity:    50,
	},
	{
		ID:          2,
		Title:       "Domain-Driven Design",
		Description: "Tackling Complexity in the Heart of Software",
		Author:      "Eric Evans",
		Quantity:    25,
	},
	{
		ID:          3,
		Title:       "Arquitetura Limpa na Prática",
		Description: "Estruture sua aplicação com profissionalismo!",
		Author:      "Otávio Lemos",
		Quantity:    10,
	},
	{
		ID:          4,
		Title:       "Clean Architecture",
		Description: "Practical Software Architecture Solutions from the Legendary Robert C. Martin (“Uncle Bob”)",
		Author:      "Robert Cecil Martin",
		Quantity:    21039,
	},
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": "true"})
	})

	app.Get("/api/healthcheck", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Get("/api/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	app.Get("/api/books/list", func(c *fiber.Ctx) error {
		var booksName = []bookName{}

		for _, book := range books {
			booksName = append(booksName, bookName{book.ID, book.Title})
		}

		return c.JSON(booksName)
	})

	app.Get("/api/books/:id", func(c *fiber.Ctx) error {
		bookIdParam, err := c.ParamsInt("id")

		if err != nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}

		for _, book := range books {
			if book.ID == bookIdParam {
				return c.JSON(book)
			}
		}

		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(BASE_URL))
	fmt.Printf("✅ Application running on %s\n", BASE_URL)
}

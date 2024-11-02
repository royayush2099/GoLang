package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2" // Fiber as an alternative to Express
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Server started...")
	app := fiber.New() // Creating an instance of the app
	todos := []Todo{}

	// Home route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello, World!"})
	})

	// Route for adding a new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := new(Todo)

		// Parse the body into the `todo` struct
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request"})
		}

		// Check if the todo body is empty
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		// Assign an ID and add to the todos list
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		// Return the newly created todo
		return c.Status(201).JSON(todo)
	})

	// Start the server on port 4000
	log.Fatal(app.Listen(":4000"))
}

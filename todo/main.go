package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// struct defines schema for todo item
type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Welcome")
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not load env vars")
	}

	PORT := os.Getenv("PORT")

	// in memory storage; empty slice of Todo structs to hold todo items
	todos := []Todo{}

	// get request, retrive all
	app.Get("/v1/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// post request, create new todo item to add to slice
	app.Post("/v1/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		// parse request body
		err := c.BodyParser(todo)
		if err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		// otherwise add todo, autoincrement id
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// update todo
	app.Patch("/v1/todos/:id", func(c *fiber.Ctx) error {
		// extract id parameter from url
		id := c.Params("id")

		// if string of id find matching todo id
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Entry not found"})
	})

	// delete todo
	app.Delete("/v1/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// if string of id find matching todo id
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				// all entries from the beginning to but not including i, then the rest
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(204).JSON(fiber.Map{"success": "Deleted"})
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Entry not found"})
	})

	log.Fatal(app.Listen(":" + PORT))
}

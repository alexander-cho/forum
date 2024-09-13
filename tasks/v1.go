package main

import (
	"fmt"
	// "log"
	// "os"
	// "github.com/gofiber/fiber/v2"
	// "github.com/joho/godotenv"
)

func v1() {
	fmt.Println("Welcome")
	// app := fiber.New()

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Could not load env vars")
	// }

	// PORT := os.Getenv("PORT")

	// // in memory storage; empty slice of Task structs to hold task items
	// tasks := []Task{}

	// // get request, retrive all
	// app.Get("/v1/tasks", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(tasks)
	// })

	// // post request, create new task item to add to slice
	// app.Post("/v1/tasks", func(c *fiber.Ctx) error {
	// 	task := &Task{}

	// 	// parse request body
	// 	err := c.BodyParser(task)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if task.Body == "" {
	// 		return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
	// 	}

	// 	// otherwise add task, autoincrement id
	// 	task.ID = len(tasks) + 1
	// 	tasks = append(tasks, *task)

	// 	return c.Status(201).JSON(task)
	// })

	// // update task
	// app.Patch("/v1/tasks/:id", func(c *fiber.Ctx) error {
	// 	// extract id parameter from url
	// 	id := c.Params("id")

	// 	// if string of id find matching task id
	// 	for i, task := range tasks {
	// 		if fmt.Sprint(task.ID) == id {
	// 			tasks[i].Completed = true
	// 			return c.Status(200).JSON(tasks[i])
	// 		}
	// 	}

	// 	return c.Status(404).JSON(fiber.Map{"error": "Entry not found"})
	// })

	// // delete task
	// app.Delete("/v1/tasks/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	// if string of id find matching task id
	// 	for i, task := range tasks {
	// 		if fmt.Sprint(task.ID) == id {
	// 			// all entries from the beginning to but not including i, then the rest
	// 			tasks = append(tasks[:i], tasks[i+1:]...)
	// 			return c.Status(204).JSON(fiber.Map{"success": "Deleted"})
	// 		}
	// 	}

	// 	return c.Status(404).JSON(fiber.Map{"error": "Entry not found"})
	// })

	// log.Fatal(app.Listen(":" + PORT))
}

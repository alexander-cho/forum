package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// struct defines schema for task item
type Task struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello from main")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env vars", err)
	}

	mongoDbUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoDbUri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB Atlas")

	collection = client.Database("go-react-tasks").Collection("tasks")

	app := fiber.New()

	app.Get("/v2/tasks", getTasks)
	app.Post("/v2/tasks", createTask)
	app.Patch("/v2/tasks/:id", updateTask)
	app.Delete("/v2/tasks/:id", deleteTask)

	PORT := os.Getenv("PORT")
	log.Fatal(app.Listen("0.0.0.0:" + PORT))

}

func getTasks(c *fiber.Ctx) error {
	var tasks []Task

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	// postpone execution of function until surrounding one completes
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			return err
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func createTask(c *fiber.Ctx) error {
	task := new(Task)

	if err := c.BodyParser(task); err != nil {
		return err
	}

	if task.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "body cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return err
	}

	task.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(task)
}

func updateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid object ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"success": true})
}

func deleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid object ID"})
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"success": true})
}

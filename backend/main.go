package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	app.Use(cors.New())

	// MongoDB connection

	app.Static("/", "../client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "usarios desde el backend",
		})
	})

	app.Listen(fmt.Sprintf(":%s", port))
	fmt.Println("Server on port 3000")
}

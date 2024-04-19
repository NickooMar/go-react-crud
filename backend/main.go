package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NickooMar/go-react-crud/database"
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
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect()
	fmt.Println("Conexión a MongoDB exitosa")

	app.Static("/", "../client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "usarios desde el backend",
		})
	})

	app.Listen(fmt.Sprintf(":%s", port))
	fmt.Println("Server on port 3000")
}

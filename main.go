package main

import (
	"github.com/Deepanshu276/AccuKnox-Task/database"
	"github.com/Deepanshu276/AccuKnox-Task/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	router.Setup(app)

	app.Listen(":8000")
}
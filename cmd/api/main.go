package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello From Web Store Ecomeerce",
		})
	})

	log.Fatal(app.Listen(":3000"))
	log.Panicln("Server is running on htpp//localhost:3000")
}

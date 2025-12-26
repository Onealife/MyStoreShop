package main

import (
	"fmt"
	"log"

	"github.com/Onealife/MyStoreShop/config"
	"github.com/Onealife/MyStoreShop/internal/adapters/db"
	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.LoadConfig()

	app := fiber.New()

	db.Connect(cfg)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello From Web Store Ecomeerce",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.AppPort)))
	log.Panicln("Server is running on htpp//localhost:", cfg.AppPort)
}

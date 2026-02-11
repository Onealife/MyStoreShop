<<<<<<< HEAD
// @title API E-commerce
// @version 1.0
// @description Authentication API with Role-based Access Control
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"

	_ "github.com/Onealife/Nutchapholshop/docs"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/Onealife/Nutchapholshop/internal/adapters/http/handlers"
	"github.com/Onealife/Nutchapholshop/internal/adapters/http/routes"
	"github.com/Onealife/Nutchapholshop/internal/adapters/persistence/repositories"
	"github.com/Onealife/Nutchapholshop/internal/config"
	"github.com/Onealife/Nutchapholshop/internal/core/domain/services"
=======
package main

import (
	"fmt"
	"log"

	"github.com/Onealife/MyStoreShop/config"
	"github.com/Onealife/MyStoreShop/internal/adapters/db"
>>>>>>> b92c0c6558e709b26467f1eda67ef5d2ddb33796
	"github.com/gofiber/fiber/v2"
)

func main() {

<<<<<<< HEAD
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load configuration %v", err)
	}

	db := config.SetupDatabase(cfg)

	userRepo := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(logger.New())
	app.Use(cors.New())

	routes.SetupRoute(app, authHandler)

	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
=======
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
>>>>>>> b92c0c6558e709b26467f1eda67ef5d2ddb33796
}

package routes

import (
	"github.com/Onealife/Nutchapholshop/internal/adapters/http/handlers"
	"github.com/Onealife/Nutchapholshop/internal/adapters/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoute(app *fiber.App, authHandler *handlers.AuthHandler) {

	adminHandler := handlers.NewAdminHandler()

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	user.Get("/profile", authHandler.GetProfile)

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RequiredRole("admin"))
	admin.Get("/dashboard", adminHandler.GetDashboard)
	admin.Post("/register", authHandler.AdminRegister)

}

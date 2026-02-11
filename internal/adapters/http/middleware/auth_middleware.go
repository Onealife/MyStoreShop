package middleware

import (
	"strings"

	"github.com/Onealife/Nutchapholshop/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	// Return the middleware function
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header required",
			})
		}

		// token format is "[1]Bearer [2]<token>" = len 2
		// tokenParts[0] = "Bearer" index 0
		// tokenParts[1] = "<token>" index 1
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}
		// Validate the token
		token := tokenParts[1]
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Store user information in context
		c.Locals("userID", claims.UserID)
		c.Locals("role", claims.Role)

		// Proceed to the next middleware/handler
		return c.Next()
	}
}

// func สำหรับตรวจสอบ role ที่ต้องการ
func RequiredRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("role").(string)

		for _, role := range roles {
			if userRole == role {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "insufficient permissions",
		})
	}
}

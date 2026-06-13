package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user/go-user-api/internal/handler"
)

func RegisterUserRoutes(router fiber.Router, userHandler *handler.UserHandler) {
	users := router.Group("/users")
	users.Post("/", userHandler.CreateUser)
	users.Get("/", userHandler.ListUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}

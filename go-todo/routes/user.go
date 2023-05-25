package routes

import (
	"github.com/labstack/echo/v4"

	databases "go-todo/database"
	"go-todo/handlers"
	"go-todo/pkg/middleware"
	"go-todo/repositories"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(databases.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	e.PATCH("/user", middleware.Auth(h.UpdateUser))
}

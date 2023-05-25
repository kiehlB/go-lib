package routes

import (
	databases "go-todo/database"
	"go-todo/handlers"
	"go-todo/pkg/middleware"
	"go-todo/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(databases.DB)
	profileRepository := repositories.RepositoryProfile(databases.DB)
	h := handlers.HandlerAuth(authRepository, profileRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}

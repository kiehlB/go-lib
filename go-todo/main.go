package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	databases "go-todo/database"
	"go-todo/models"
	"go-todo/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	databases.InitPGDatabase(
		&models.User{},
		&models.Post{},
		&models.Profile{},
	)

	routes.RouteInit(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start(":4000"))
}

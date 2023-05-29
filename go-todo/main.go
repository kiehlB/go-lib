package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	databases "go-todo/database"
	"go-todo/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	databases.InitPGDatabase()

	routes.RouteInit(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start(":4000"))
}

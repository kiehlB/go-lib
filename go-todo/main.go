package main

import (
	"log"

	databases "go-todo/database"
	"go-todo/entities"
	"go-todo/modules/users"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if dotEnvError := godotenv.Load(); dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}

	databases.InitPGDatabase(
		&entities.User{},
	)
}

func main() {
	e := echo.New()
 
	api := e.Group("/v1")

	users.UserRoutes(api)


	e.Logger.Fatal(e.Start(":1323"))
}
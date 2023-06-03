package routes

import (
	databases "go-todo/database"
	"go-todo/handlers"
	"go-todo/repositories"

	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Group) {
	postRepository := repositories.RepositoryPost(databases.DB)
	h := handlers.HandlerPost(postRepository)

	e.GET("/post", h.FindPosts)
	e.POST("/create", h.CreatePost)
	e.POST("/update", h.UpdatePost)
	e.DELETE("/delete", h.DeletePost)

}

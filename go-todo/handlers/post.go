package handlers

import (
	postsdto "go-todo/dto/post"
	"go-todo/models"
	"go-todo/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handlerPost struct {
	PostRepository repositories.PostRepository
}

func HandlerPost(PostRepository repositories.PostRepository) *handlerPost {
	return &handlerPost{
		PostRepository: PostRepository,
	}
}

func (h *handlerPost) FindPosts(c echo.Context) error {
	posts, err := h.PostRepository.GetAllPosts()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch posts",
		})
	}

	return c.JSON(http.StatusOK, posts)

}

func (h *handlerPost) CreatePost(c echo.Context) error {
	var createDTO postsdto.PostCreateRequest
	if err := c.Bind(&createDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	post := models.Post{
		Content: createDTO.Content,
		UserID:  createDTO.UserID,
	}

	createdPost, err := h.PostRepository.CreatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create post",
		})
	}

	return c.JSON(http.StatusCreated, createdPost)
}

func (h *handlerPost) UpdatePost(c echo.Context) error {

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid post ID",
		})
	}

	var updateDTO postsdto.PostUpdateRequest
	if err := c.Bind(&updateDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	post, err := h.PostRepository.GetPostByID(postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to get post",
		})
	}
	if post == nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Post not found",
		})
	}

	post.Content = updateDTO.Content

	err = h.PostRepository.UpdatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to update post",
		})
	}

	return c.JSON(http.StatusOK, post)
}

func (h *handlerPost) DeletePost(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid post ID",
		})
	}

	post, err := h.PostRepository.GetPostByID(postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to get post",
		})
	}
	if post == nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Post not found",
		})
	}

	err = h.PostRepository.DeletePost(postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to delete post",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Post deleted successfully",
	})
}

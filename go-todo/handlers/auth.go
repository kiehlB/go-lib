package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	dto "go-todo/dto"
	authdto "go-todo/dto/auth"
	"go-todo/models"
	"go-todo/pkg/bcrypt"
	jwtToken "go-todo/pkg/jwt"
	"go-todo/repositories"

	"log"
	"net/http"
	"time"
)

type handlerAuth struct {
	AuthRepository    repositories.AuthRepository
	ProfileRepository repositories.ProfileRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository, ProfileRepository repositories.ProfileRepository) *handlerAuth {
	return &handlerAuth{
		AuthRepository:    AuthRepository,
		ProfileRepository: ProfileRepository,
	}
}

func (h *handlerAuth) Register(c echo.Context) error {

	request := new(authdto.AuthRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	_, err = h.AuthRepository.CheckEmailValid(user.Email)
	if err != nil {
		data, err := h.AuthRepository.Register(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, dto.SuccessResponse{Status: http.StatusOK, Message: "Your registration is successful", Data: data})
	} else {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: http.StatusBadRequest, Message: "This email is already registered"})
	}
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.CheckEmailValid(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: http.StatusBadRequest, Message: "wrong password"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID

	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	_, err = h.ProfileRepository.GetProfile(user.ID)
	if err != nil {
		profile := models.Profile{
			ID:     user.ID,
			UserID: user.ID,
		}
		_, err = h.ProfileRepository.CreateProfile(profile)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		}
	}

	loginResponse := authdto.LoginResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: isValid,
		Token:    token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{Status: http.StatusOK, Message: "You have successfully logged in", Data: loginResponse})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))

	return c.JSON(http.StatusOK, dto.SuccessResponse{Status: http.StatusOK, Message: "The authentication was successfully examined", Data: user})
}

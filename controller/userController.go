package controller

import (
	"masaku/helpers"
	"masaku/models"
	"masaku/models/web"
	"masaku/utils"
	"masaku/utils/req"
	"masaku/utils/res"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UsersControlInterface interface {
	Show(c echo.Context) error
	Register(c echo.Context) error
	LoginUser(c echo.Context) error
	Profile(c echo.Context) error
}

type UsersModel struct {
	db *gorm.DB
	jwt helpers.JWTInterface
}

func NewUsersControl(db *gorm.DB, jwt helpers.JWTInterface) UsersControlInterface {
	return &UsersModel{db: db, jwt : jwt}
}

// Show retrieves a user by ID
func (um *UsersModel) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var user models.User
	if err := um.db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found"))
	}

	response := res.ConvertGeneral(&user)
	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", response))
}


func (um *UsersModel) Register(c echo.Context) error {
	var user web.UserRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	// Validate user input
	if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Email and Password are required"))
	}

	userDb := req.PassBody(user)

	// Hash the password
	userDb.Password = helpers.HashPassword(userDb.Password)

	if err := um.db.Create(&userDb).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store user data"))
	}

	response := res.ConvertGeneral(userDb)
	return c.JSON(http.StatusCreated, utils.SuccessResponse("User successfully created", response))
}

// LoginUser handles user login
func (um *UsersModel) LoginUser(c echo.Context) error {
	var loginRequest web.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var user models.User
	if err := um.db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	if err := helpers.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	token := um.jwt.GenerateJWT(uint(user.ID), user.Nama)

	response := web.UserLoginResponse{
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Login successful", response))
}

// Profile retrieves the profile of the authenticated user
func (um *UsersModel) Profile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := int(claims["id"].(float64))

	var profile models.User
	if err := um.db.First(&profile, ID).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found"))
	}

	response := res.ConvertGeneral(&profile)
	return c.JSON(http.StatusOK, utils.SuccessResponse("User profile retrieved successfully", response))
}

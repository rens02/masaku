package routes

import (
	"masaku/config"
	"masaku/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// RouteUser defines routes for user-related actions
func RouteUser(e *echo.Echo, uc controller.UsersControlInterface, cfg config.ProgramConfig) {
	// Public routes
	e.POST("/users", uc.Register)       // Register a new user
	e.POST("/users/login", uc.LoginUser) // User login

	// Protected routes
	protected := e.Group("/user", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.Secret),
	}))

	protected.GET("/profile", uc.Profile)            // Get authenticated user's profile
	protected.GET("/:id", uc.Show)    // Update authenticated user's profile
}

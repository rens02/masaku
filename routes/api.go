package routes

import (
	"masaku/config"
	"masaku/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// RouteUser defines routes for user-related actions
func RouteUser(e *echo.Echo, uc controller.UsersControlInterface, rc controller.ResepControlInterface, kc controller.KategoriControlInterface, cfg config.ProgramConfig) {
	// Public routes
	e.POST("masaku/users", uc.Register)       // Register a new user
	e.POST("masaku/users/login", uc.LoginUser) // User login

	// Protected routes
	protected := e.Group("masaku/", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.Secret),
	}))

	protected.GET("users/profile", uc.Profile)            // Get authenticated user's profile
	protected.GET("users/:id", uc.Show)    // Update authenticated user's profile

	protected.GET("reseps/:id", rc.ShowResep)
	protected.GET("reseps", rc.ShowAllResep)

	protected.GET("kategori/:id", kc.ShowKategori)
	protected.GET("kategori", kc.ShowAllKategori)

}

package controller

import (
	"masaku/models"
	"masaku/utils"
	"masaku/utils/res"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ResepControlInterface interface {
	ShowResep(c echo.Context) error
	ShowAllResep(c echo.Context) error
}

type ResepModel struct {
	db *gorm.DB
	
}

func NewResepControl(db *gorm.DB) ResepControlInterface {
	return &UsersModel{db: db}
}

// Show retrieves a resep by ID
func (um *UsersModel) ShowResep(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var resep models.Resep
	if err := um.db.First(&resep, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("resep not found"))
	}

	response := res.ConvertGeneralResep(&resep)
	return c.JSON(http.StatusOK, utils.SuccessResponse("resep data successfully retrieved", response))
}

// Show retrieves a resep 
func (um *UsersModel) ShowAllResep(c echo.Context) error {
	
	var reseps []models.Resep
	if err := um.db.Select(&reseps).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("resep not found"))
	}

	response := res.ConvertResep(reseps)
	return c.JSON(http.StatusOK, utils.SuccessResponse("resep data successfully retrieved", response))
}

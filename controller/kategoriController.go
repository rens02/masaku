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

type KategoriControlInterface interface {
	ShowKategori(c echo.Context) error
	ShowAllKategori(c echo.Context) error
}

type KategoriModel struct {
	db *gorm.DB
	
}

func NewKategoriControl(db *gorm.DB) KategoriControlInterface {
	return &UsersModel{db: db}
}

// Show retrieves a Kategori by ID
func (um *UsersModel) ShowKategori(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var kategori models.Kategori
	if err := um.db.First(&kategori, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Kategori not found"))
	}

	response := res.ConvertGeneralKategori(&kategori)
	return c.JSON(http.StatusOK, utils.SuccessResponse("Kategori data successfully retrieved", response))
}


// Show retrieves a Kategori 
func (um *UsersModel) ShowAllKategori(c echo.Context) error {
	
	var Kategoris []models.Kategori
	if err := um.db.Select(&Kategoris).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Kategori not found"))
	}

	response := res.ConvertKategori(Kategoris)
	return c.JSON(http.StatusOK, utils.SuccessResponse("Kategori data successfully retrieved", response))
}

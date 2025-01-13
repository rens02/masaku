package controller

import (
	"masaku/models"
	"masaku/models/web"
	"masaku/utils"
	"masaku/utils/res"
	"net/http"
	"masaku/helpers"
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GenerateControllInterface interface {
	Generate(c echo.Context) error
}

type GenerateModel struct {
	db *gorm.DB
	jwt helpers.JWTInterface
	openAi helpers.OpenAiInterface
}


func NewGenerateControl(db *gorm.DB, jwt helpers.JWTInterface, openAi helpers.OpenAiInterface ) GenerateControllInterface {
	return &GenerateModel{db: db, jwt : jwt, openAi : openAi}
}

func (ug *GenerateModel) Generate(c echo.Context) error {

	var token = c.Get("user")
	id := ug.jwt.ExtractToken(token.(*jwt.Token))

	var user models.User
	if err := ug.db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found"))
	}

	var generate web.GenerateRequest

	if err := c.Bind(&generate); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	saran := ug.openAi.GenerateSaran(generate.Maag, generate.Asam_urat, generate.Hipertensi)

	user.Maag = generate.Maag
	user.Asam_urat = generate.Asam_urat
	user.Hipertensi = generate.Hipertensi
	user.Saran = saran.Saran


	var qry = ug.db.Save(&user)
	if err := qry.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store user data"))
	}

	response := res.ConvertGeneral(&user)
	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", response))
}
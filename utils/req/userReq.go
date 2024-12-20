package req

import (
	"masaku/models"
	"masaku/models/web"
)

func PassBody(users web.UserRequest) *models.User {
	return &models.User{
		Nama:     users.Nama,
		Email:    users.Email,
		Password: users.Password,
	}
}

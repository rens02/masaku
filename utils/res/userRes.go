package res

import (
	"masaku/models"
	"masaku/models/web"
)

func ConvertIndex(users []models.User) []web.UserReponse {
	var results []web.UserReponse
	for _, user := range users {
		userResponse := web.UserReponse{
			Id:         int(user.ID),
			Nama:       user.Nama,
			Email:      user.Email,
			Foto:       user.Foto,
			Diabetes:   user.Diabetes,
			Maag:       user.Maag,
			Asam_urat:  user.Asam_urat,
			Hipertensi: user.Hipertensi,
		}
		results = append(results, userResponse)
	}

	return results
}

func ConvertGeneral(user *models.User) web.UserReponse {
	return web.UserReponse{
		Id:         int(user.ID),
		Nama:       user.Nama,
		Email:      user.Email,
		Foto:       user.Foto,
		Diabetes:   user.Diabetes,
		Maag:       user.Maag,
		Asam_urat:  user.Asam_urat,
		Hipertensi: user.Hipertensi,
	}
}

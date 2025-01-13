package req

import (
	"masaku/models"
	"masaku/models/web"
)

func PassKategoriBody(kategori web.KategoriRequest) *models.Kategori {
	return &models.Kategori{
		Nama_Kategori: kategori.Nama_Kategori,
		Foto:          kategori.Foto,
	}
}

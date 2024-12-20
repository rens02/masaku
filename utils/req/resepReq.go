package req

import (
	"masaku/models"
	"masaku/models/web"
)

func PassResepBody(resep web.ResepRequest) *models.Resep {
	return &models.Resep{
		Nama_Makanan:   resep.Nama_Makanan,
		Bahan:          resep.Bahan,
		Cara_Pembuatan: resep.Cara_Pembuatan,
		Foto:           resep.Foto,
		KategoriID:     resep.KategoriID,
	}
}

package res

import (
	"masaku/models"
	"masaku/models/web"
)

func ConvertResep(reseps []models.Resep) []web.ResepResponse {
	var results []web.ResepResponse
	for _, resep := range reseps {
		resepResponse := web.ResepResponse{
			ID:             int(resep.ID),
			Nama_Makanan:   resep.Nama_Makanan,
			Bahan:          resep.Bahan,
			Cara_Pembuatan: resep.Cara_Pembuatan,
			Foto:           resep.Foto,
			KategoriID:     resep.KategoriID,
		}
		results = append(results, resepResponse)
	}

	return results
}

func ConvertGeneralResep(resep *models.Resep) web.ResepResponse {
	return web.ResepResponse{
		ID:             int(resep.ID),
		Nama_Makanan:   resep.Nama_Makanan,
		Bahan:          resep.Bahan,
		Cara_Pembuatan: resep.Cara_Pembuatan,
		Foto:           resep.Foto,
		KategoriID:     resep.KategoriID,
	}
}

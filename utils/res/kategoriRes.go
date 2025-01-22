package res

import (
	"masaku/models"
	"masaku/models/web"
)

func ConvertKategori(kategoris []models.Kategori) []web.KategoriResponse {
	var results []web.KategoriResponse
	for _, kategori := range kategoris {
		kategoriResponse := web.KategoriResponse{
			ID:            int(kategori.ID),
			Nama_Kategori: kategori.Nama_Kategori,
		}
		results = append(results, kategoriResponse)
	}

	return results
}

func ConvertGeneralKategori(kategori *models.Kategori) web.KategoriResponse {
	return web.KategoriResponse{
		ID:            int(kategori.ID),
		Nama_Kategori: kategori.Nama_Kategori,
	}
}

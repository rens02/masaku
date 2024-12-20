package web

type ResepResponse struct {
	ID             int    `json:"id"`
	Nama_Makanan   string `json:"nama_makanan"`
	Bahan          string `json:"bahan"`
	Cara_Pembuatan string `json:"cara_pembuatan"`
	Foto           string `json:"foto"`
	KategoriID     uint   `json:"kategori_id"`
	Kategori       string `json:"kategori"`
}

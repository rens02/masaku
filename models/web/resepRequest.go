package web

type ResepRequest struct {
	Nama_Makanan   string `json:"nama_makanan" form:"nama_makanan"`
	Bahan          string `json:"bahan" form:"bahan"`
	Cara_Pembuatan string `json:"cara_pembuatan" form:"cara_pembuatan"`
	Foto           string `json:"foto" form:"foto"`
	KategoriID     uint   `json:"kategori_id" form:"kategori_id"`
	Kategori       string `json:"kategori" form:"kategori"`
}

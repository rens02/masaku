package web

type KategoriRequest struct {
	Nama_Kategori string `json:"nama_kategori" form:"nama_kategori"`
	Foto          string `json:"foto" form:"foto"`
}
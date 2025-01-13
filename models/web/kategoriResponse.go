package web

type KategoriResponse struct {
	ID            int    `json:"id"`
	Nama_Kategori string `json:"nama_kategori"`
	Foto          string `json:"foto"`
}

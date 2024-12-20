package web

type UserReponse struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama" form:"nama"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Foto       string `json:"foto" form:"foto"`
	Diabetes   int    `json:"diabetes" form:"diabetes"`
	Maag       int    `json:"maag" form:"maag"`
	Asam_urat  int    `json:"asam_urat" form:"asam_urat"`
	Hipertensi int    `json:"hipertensi" form:"hipertensi"`
}

type UserLoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

package web

type GenerateRequest struct {
	Maag       int    `json:"maag" form:"maag"`
	Asam_urat  int    `json:"asam_urat" form:"asam_urat"`
	Hipertensi int    `json:"hipertensi" form:"hipertensi"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Resep struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Nama_Makanan   string
	Bahan          string
	Cara_Pembuatan string
	Foto           string
	KategoriID     uint
	Kategori       Kategori `gorm:"ForeignKey:KategoriID"`
}

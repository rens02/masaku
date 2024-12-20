package models

import (
	"time"

	"gorm.io/gorm"
)

type Kategori struct {
	ID        		uint `gorm:"primarykey"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Nama_Kategori   string
	Foto	  		string
}
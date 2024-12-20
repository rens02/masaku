package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Name       string
	Email      string
	Password   string
	Foto       string
	Diabetes   int
	Maag       int
	Asam_urat  int
	Hipertensi int
}

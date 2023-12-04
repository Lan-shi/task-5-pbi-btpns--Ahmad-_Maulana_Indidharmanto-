package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Path     string `gorm:"type:varchar(255)"`
	ID       uint   `gorm:"primary_key"`
	Title    string `gorm:"type:varchar(100);not null"`
	Caption  string `gorm:"type:varchar(200);not null"`
	PhotoUrl string `gorm:"type:varchar(200);not null"`
	UserID   uint
	User     User `gorm:"foreignKey:UserID"`
}

package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name  string `gorm:"not null;unique_index" json:"name"`
	Price int    `gorm:"not null" json:"price"`
	Img   string `gorm:"default:https://st3.depositphotos.com/1419868/13478/i/950/depositphotos_134786356-stock-photo-burger-and-french-fries.jpg" json:"img"`
	
}

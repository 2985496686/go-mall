package model

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	ShopName string
}

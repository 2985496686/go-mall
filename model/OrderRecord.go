package model

import "gorm.io/gorm"

type OrderRecord struct {
	gorm.Model
	UserEmail string `json:"userEmail"`
	Status    int    `json:"status" default:"0"`
}

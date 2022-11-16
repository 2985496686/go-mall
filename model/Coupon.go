package model

import (
	"gorm.io/gorm"
	"time"
)

type Coupon struct {
	gorm.Model
	Title     string    `json:"title" gorm:"not null  type:varchar(25)"`
	PayMoney  uint      `json:"payMoney" gorm:"not null"`
	StartTime time.Time `json:"startTime" gorm:"not null"`
	EndTime   time.Time `json:"endTime" gorm:"not null "`
	Stock     uint      `json:"stock" gorm:"not null"`
}

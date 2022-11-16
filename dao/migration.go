package dao

import (
	"fmt"
	"mall/model"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&model.Address{},
		&model.User{},
		&model.Coupon{},
		&model.OrderRecord{},
	)
	if err != nil {
		panic(fmt.Sprintf("dataBase create failed:%s", err.Error()))
	}
}

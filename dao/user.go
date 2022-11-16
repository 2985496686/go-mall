package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(crx context.Context) *UserDao {
	return &UserDao{
		NewDBClient(crx),
	}
}

func (dao *UserDao) ExistOrNotByEmail(email string) (user model.User, exit bool, err error) {
	var count int64
	err = dao.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Find(&user).Error
	if count == 0 || err != nil {
		return user, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.db.Create(user).Error
}

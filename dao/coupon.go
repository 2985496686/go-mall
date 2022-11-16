package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/dto/result"
	"mall/model"
)

type CouponDao struct {
	db *gorm.DB
}

func NewCouponDao(ctx context.Context) *CouponDao {
	return &CouponDao{
		db: NewDBClient(ctx),
	}
}

func (dao *CouponDao) CouponList() *result.Result {
	var couponList []model.Coupon
	if err := dao.db.Model(&model.Coupon{}).Order("start_time").Find(&couponList).Error; err != nil {
		return result.Fail(err.Error())
	}
	return result.Ok(couponList)
}

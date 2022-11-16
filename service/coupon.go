package service

import (
	"github.com/gin-gonic/gin"
	"mall/dao"
	"mall/dto/result"
)

type CouponService struct {
}

func (service *CouponService) CouponList(ctx *gin.Context) *result.Result {
	couponDao := dao.NewCouponDao(ctx.Request.Context())
	return couponDao.CouponList()
}

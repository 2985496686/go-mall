package v1

import (
	"github.com/gin-gonic/gin"
	"mall/service"
	"net/http"
)

func CouponList(ctx *gin.Context) {
	couponService := service.CouponService{}
	result := couponService.CouponList(ctx)
	ctx.JSON(http.StatusOK, result)
}

package routes

import (
	"github.com/gin-gonic/gin"
	v1 "mall/api/v1"
	"mall/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//v1
	v1Group := r.Group("/api/v1", middleware.RefreshToken())
	{
		userGroup := v1Group.Group("/user")
		{
			userGroup.POST("/code", v1.SendCode)
			userGroup.POST("/login", v1.Login)
			userGroup.GET("/info:id", v1.Info)
			userGroup.GET("/me", middleware.VerifyLogin(), v1.GetMe)
		}

		shopGroup := v1Group.Group("/coupon")
		{
			shopGroup.POST("/create", v1.CouponList)
			shopGroup.GET("/list")
			//shopGroup.POST("/seckill", v1.Seckill)
		}
	}
	return r
}

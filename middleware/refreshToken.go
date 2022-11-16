package middleware

import (
	"github.com/gin-gonic/gin"
	"mall/cache"
	"mall/service"
	"time"
)

func RefreshToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		if token := context.GetHeader("authorization"); token != "" {
			//刷新token
			c := cache.NewRedisClient()
			c.Expire(service.LoginUser+token, time.Second*service.LoginUserDdl)
		}
		return
	}
}

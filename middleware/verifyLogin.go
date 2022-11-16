package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"mall/cache"
	"mall/dto"
	"mall/dto/result"
	"mall/service"
	"net/http"
)

// VerifyLogin 验证用户的登录状态
func VerifyLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取token
		token := context.GetHeader("authorization")
		if token == "" {
			context.JSON(http.StatusOK, result.Fail("用户未登录!"))
			context.Abort()
			return
		}
		//获取用户信息
		c := cache.NewRedisClient()
		uJson, err := c.Get(service.LoginUser + token).Result()
		if err != nil {
			context.JSON(http.StatusOK, result.Fail(err.Error()))
			context.Abort()
			return
		}
		user := dto.UserDto{}
		if err = json.Unmarshal([]byte(uJson), &user); err != nil {
			context.JSON(http.StatusOK, result.Fail(err.Error()))
			context.Abort()
			return
		}
		context.Set("user", user)
	}
}

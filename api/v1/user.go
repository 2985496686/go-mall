package v1

import (
	"github.com/gin-gonic/gin"
	"mall/dto/result"
	"mall/service"
	"net/http"
)

func SendCode(ctx *gin.Context) {
	email := ctx.PostForm("email")
	userService := service.UserService{}
	r := *userService.SendCode(email)
	ctx.JSON(http.StatusOK, r)
}

func Info(ctx *gin.Context) {

}

func Login(ctx *gin.Context) {
	userService := service.UserService{}
	if err := ctx.ShouldBind(&userService); err == nil {
		r := *userService.Login(ctx.Request.Context())
		ctx.JSON(http.StatusOK, r)
		return
	} else {
		ctx.JSON(http.StatusBadRequest, "错误的参数")
		return
	}
}

func GetMe(ctx *gin.Context) {
	user := ctx.Value("user")
	ctx.JSON(http.StatusOK, result.Ok(user))
	return
}

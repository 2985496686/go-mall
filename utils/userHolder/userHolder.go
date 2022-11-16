package userHolder

import (
	"context"
	"mall/dto"
)

var (
	ctx = context.TODO()
)

func SaveUser(user dto.UserDto) {
	context.WithValue(ctx, "user", user)
}

func GetUser() dto.UserDto {
	user := ctx.Value("user")
	return user.(dto.UserDto)
}

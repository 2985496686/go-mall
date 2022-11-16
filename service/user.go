package service

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"mall/cache"
	"mall/dao"
	"mall/dto"
	"mall/dto/result"
	"mall/model"
	"mall/utils"
	"time"
)

type UserService struct {
	Code     string `json:"code" form:"code"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

func (service *UserService) SendCode(email string) *result.Result {
	//检验邮箱格式
	if ok := utils.VerifyEmailFormat(email); !ok {
		return result.Fail("错误的邮箱！")
	}
	code := utils.RandomCode()
	//将验证码保存到redis
	c := cache.NewRedisClient()
	//60s内禁止重复发送验证码
	if cnt := c.Exists(LoginCodeKey + email).Val(); cnt == 1 {
		return result.Fail("禁止重复发送验证码！")
	}
	c.Set(LoginCodeKey+email, code, time.Second*LoginCodeDdl)
	return result.Ok(nil)
}

func (service *UserService) Login(ctx context.Context) *result.Result {
	c := cache.NewRedisClient()
	userDao := dao.NewUserDao(ctx)
	//校验邮箱
	if ok := utils.VerifyEmailFormat(service.Email); !ok {
		return result.Fail("错误的邮箱!")
	}
	//校验验证码
	code, err := c.Get(LoginCodeKey + service.Email).Result()
	if err != nil {
		return result.Fail(err.Error())
	}
	if code != service.Code {
		return result.Fail("验证码错误!")
	}
	//生成随机token
	token := uuid.NewString()
	//检验当前用户是否已经注册
	user, exit, err := userDao.ExistOrNotByEmail(service.Email)
	if err != nil {
		return result.Fail(err.Error())
	}
	//用户不存在，创建对象
	if !exit {
		user = model.User{
			NickName: UserNamePrefix + uuid.NewString()[0:10],
			Email:    service.Email,
			Password: "",
		}
		err = userDao.CreateUser(&user)
	}
	if err != nil {
		result.Fail(err.Error())
	}
	//将用户保存基本信息保存到redis
	userDto := dto.UserDto{
		NickName: user.NickName,
		Email:    user.Email,
	}
	uJson, _ := json.Marshal(userDto)
	err = c.Set(LoginUser+token, uJson, time.Second*LoginUserDdl).Err()
	if err != nil {
		return result.Fail(err.Error())
	}
	return result.Ok(token)
}

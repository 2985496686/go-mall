package dto

type UserDto struct {
	NickName string `json:"nickName" toMap:"nickName"`
	Email    string `json:"email" toMap:"email"`
}

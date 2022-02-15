package dto

import "web/model"

type UserDto struct {
	Name string `json:"name"`
	Tel  string `json:"tel"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
		Tel:  user.Tel,
	}
}

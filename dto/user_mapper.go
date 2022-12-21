package dto

import (
	"freeDayManager/entity"
)

func RegisterToUser(register UserRequest) entity.User {
	return entity.User{Username: register.Username, FullName: register.FullName, Password: register.Password, Email: register.Email}
}

func UserToDto(user entity.User) UserDto {
	return UserDto{
		Id:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Roles:    ToArray(user.Roles)}
}

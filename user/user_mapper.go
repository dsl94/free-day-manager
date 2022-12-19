package user

import (
	"freeDayManager/role"
)

func RegisterToUser(register UserRequest) User {
	return User{Username: register.Username, FullName: register.FullName, Password: register.Password, Email: register.Email}
}

func UserToDto(user User) UserDto {
	return UserDto{
		Id:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Roles:    role.ToArray(user.Roles)}
}

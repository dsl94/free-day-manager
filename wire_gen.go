// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"freeDayManager/role"
	"freeDayManager/user"
	"github.com/jinzhu/gorm"
)

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Injectors from wire.go:

func InitRoleController(db *gorm.DB) role.RoleController {
	roleRepository := role.ProvideRoleRepository(db)
	roleService := role.ProvideRoleService(roleRepository)
	roleController := role.ProvideRoleController(roleService)
	return roleController
}

func InitUserController(db *gorm.DB, repository role.RoleRepository) user.UserController {
	userRepository := user.ProvideUserRepository(db)
	userService := user.ProvideUserService(userRepository, repository)
	userController := user.ProvideUserController(userService)
	return userController
}
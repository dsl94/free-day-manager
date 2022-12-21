// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"freeDayManager/controller"
	"freeDayManager/database"
	"freeDayManager/service"
	"github.com/jinzhu/gorm"
)

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Injectors from wire.go:

func InitRoleController(db *gorm.DB) controller.RoleController {
	roleRepository := database.ProvideRoleRepository(db)
	roleService := service.ProvideRoleService(roleRepository)
	roleController := controller.ProvideRoleController(roleService)
	return roleController
}

func InitUserController(db *gorm.DB, repository database.RoleRepository) controller.UserController {
	userRepository := database.ProvideUserRepository(db)
	userService := service.ProvideUserService(userRepository, repository)
	userController := controller.ProvideUserController(userService)
	return userController
}

func InitFreeDayController(db *gorm.DB) controller.FreeDayController {
	freeDayRepository := database.ProvideFreeDayRepository(db)
	freeDayService := service.ProvideFreeDayService(freeDayRepository)
	freeDayController := controller.ProvideFreeDayController(freeDayService)
	return freeDayController
}

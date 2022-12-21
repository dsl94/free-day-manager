//go:build wireinject
// +build wireinject

package main

import (
	"freeDayManager/controller"
	"freeDayManager/database"
	"freeDayManager/service"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitRoleController(db *gorm.DB) controller.RoleController {
	wire.Build(database.ProvideRoleRepository, service.ProvideRoleService, controller.ProvideRoleController)

	return controller.RoleController{}
}

func InitUserController(db *gorm.DB, repository database.RoleRepository) controller.UserController {
	wire.Build(database.ProvideUserRepository, service.ProvideUserService, controller.ProvideUserController)

	return controller.UserController{}
}

func InitFreeDayController(db *gorm.DB) controller.FreeDayController {
	wire.Build(database.ProvideFreeDayRepository, service.ProvideFreeDayService, controller.ProvideFreeDayController)

	return controller.FreeDayController{}
}

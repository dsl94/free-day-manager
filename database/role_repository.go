package database

import (
	"freeDayManager/entity"
	"github.com/jinzhu/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func ProvideRoleRepository(DB *gorm.DB) RoleRepository {
	return RoleRepository{DB: DB}
}

func (r *RoleRepository) Save(role entity.Role) entity.Role {
	r.DB.Save(&role)

	return role
}

func (r *RoleRepository) FindByRole(roleName string) entity.Role {
	var role entity.Role
	r.DB.Where(&entity.Role{RoleName: roleName}).First(&role)

	return role
}

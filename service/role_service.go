package service

import (
	"freeDayManager/database"
	"freeDayManager/entity"
)

type RoleService struct {
	RoleRepository database.RoleRepository
}

func ProvideRoleService(r database.RoleRepository) RoleService {
	return RoleService{RoleRepository: r}
}

func (r *RoleService) Save(role entity.Role) entity.Role {
	r.RoleRepository.Save(role)

	return role
}

func (r *RoleService) FindByRoleName(name string) entity.Role {
	role := r.RoleRepository.FindByRole(name)

	return role
}

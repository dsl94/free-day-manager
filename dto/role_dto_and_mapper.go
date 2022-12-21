package dto

import "freeDayManager/entity"

type RoleDto struct {
	Name string `json:"name"`
}

func ToRole(dto RoleDto) entity.Role {
	role := entity.Role{RoleName: dto.Name}

	return role
}

func ToDto(role entity.Role) RoleDto {
	dto := RoleDto{Name: role.RoleName}

	return dto
}

func ToArray(roles []entity.Role) []string {
	var res = make([]string, len(roles))
	for i, role := range roles {
		res[i] = role.RoleName
	}

	return res
}

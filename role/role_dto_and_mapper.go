package role

type RoleDto struct {
	Name string `json:"name"`
}

func ToRole(dto RoleDto) Role {
	role := Role{RoleName: dto.Name}

	return role
}

func ToDto(role Role) RoleDto {
	dto := RoleDto{Name: role.RoleName}

	return dto
}

func ToArray(roles []Role) []string {
	var res = make([]string, len(roles))
	for i, role := range roles {
		res[i] = role.RoleName
	}

	return res
}

package user

import "freeDayManager/role"

type UserService struct {
	UserRepository UserRepository
	RoleRepository role.RoleRepository
}

func ProvideUserService(u UserRepository, r role.RoleRepository) UserService {
	return UserService{UserRepository: u, RoleRepository: r}
}

func (u *UserService) FindAll() []UserDto {
	users := u.UserRepository.FindAll()
	var dtos = make([]UserDto, len(users))

	for i, user := range users {
		dtos[i] = UserToDto(user)
	}

	return dtos
}

func (u *UserService) FindById(id uint) UserDto {
	dbUser := u.UserRepository.FindById(id)

	return UserToDto(dbUser)
}

func (u *UserService) Create(user User, roleName string) User {
	dbRole := u.RoleRepository.FindByRole(roleName)

	user.Roles = []role.Role{dbRole}

	u.UserRepository.Save(user)

	return user
}

func (u *UserService) Delete(user User) {
	u.UserRepository.Delete(user)
}

func (u *UserService) FindByUsername(username string) User {
	return u.UserRepository.FindByUsername(username)
}

func (u *UserService) CheckLoginCredentials(login UserLogin) (User, bool) {
	user := u.UserRepository.FindByUsername(login.Username)

	err := user.CheckPassword(login.Password)
	if err != nil {
		return User{}, false
	}

	return user, true
}

package service

import (
	"freeDayManager/database"
	"freeDayManager/dto"
	"freeDayManager/entity"
	"log"
)

type UserService struct {
	UserRepository database.UserRepository
	RoleRepository database.RoleRepository
}

func ProvideUserService(u database.UserRepository, r database.RoleRepository) UserService {
	return UserService{UserRepository: u, RoleRepository: r}
}

func (u *UserService) FindAll() []dto.UserDto {
	users := u.UserRepository.FindAll()
	var dtos = make([]dto.UserDto, len(users))

	for i, user := range users {
		dtos[i] = dto.UserToDto(user)
	}

	return dtos
}

func (u *UserService) FindById(id uint) dto.UserDto {
	dbUser := u.UserRepository.FindById(id)
	return dto.UserToDto(dbUser)
}

func (u *UserService) Create(user entity.User, roleName string) entity.User {
	dbRole := u.RoleRepository.FindByRole(roleName)

	user.Roles = []entity.Role{dbRole}

	u.UserRepository.Save(user)

	return user
}

func (u *UserService) Update(id uint, request dto.UserRequest) dto.UserDto {
	existing := u.UserRepository.FindById(id)
	dbRole := u.RoleRepository.FindByRole(request.Role)

	existing.Roles = []entity.Role{dbRole}
	existing.FullName = request.FullName
	existing.Email = request.Email
	existing.Username = request.Username

	if request.Password != "" {
		log.Println("Changing pass")
		existing.HashPassword(request.Password)
	}

	u.UserRepository.Save(existing)

	return dto.UserToDto(existing)
}

func (u *UserService) Delete(id uint) bool {
	existing := u.UserRepository.FindById(id)
	if existing.ID > 0 {
		u.UserRepository.Delete(existing)
		return true
	} else {
		return false
	}
}

func (u *UserService) FindByUsername(username string) entity.User {
	return u.UserRepository.FindByUsername(username)
}

func (u *UserService) CheckLoginCredentials(login dto.UserLogin) (entity.User, bool) {
	user := u.UserRepository.FindByUsername(login.Username)

	err := user.CheckPassword(login.Password)
	if err != nil {
		return entity.User{}, false
	}

	return user, true
}

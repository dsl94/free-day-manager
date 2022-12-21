package database

import (
	"freeDayManager/entity"
	"github.com/jinzhu/gorm"
	"log"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) FindAll() []entity.User {
	var users []entity.User
	u.DB.Preload("Roles").Find(&users)

	return users
}

func (u *UserRepository) FindById(id uint) entity.User {
	var user entity.User
	u.DB.Preload("Roles").Find(&user, id)
	log.Println(user)
	return user
}

func (u *UserRepository) Save(user entity.User) entity.User {
	u.DB.Save(&user)

	return user
}

func (u *UserRepository) Delete(user entity.User) {
	u.DB.Delete(&user)
}

func (u *UserRepository) FindByUsername(username string) entity.User {
	var user entity.User
	u.DB.Where(&entity.User{Username: username}).Preload("Roles").First(&user)

	return user
}

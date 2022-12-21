package controller

import (
	"fmt"
	"freeDayManager/dto"
	"freeDayManager/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService service.UserService
}

func ProvideUserController(u service.UserService) UserController {
	return UserController{UserService: u}
}

func (u *UserController) Create(c *gin.Context) {
	var userRegister dto.UserRequest
	err := c.ShouldBind(&userRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	user := dto.RegisterToUser(userRegister)
	user.HashPassword(userRegister.Password)

	u.UserService.Create(user, userRegister.Role)

	c.Status(http.StatusOK)
}

func (u *UserController) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *UserController) FindOne(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		panic("Id not valid format, it must be int")
	}
	user := u.UserService.FindById(uint(id))

	if user.Id > 0 {
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}

func (u *UserController) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		panic("Id not valid format, it must be int")
	}
	var userRegister dto.UserRequest
	err2 := c.ShouldBind(&userRegister)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err2)})
		return
	}

	u.UserService.Update(uint(id), userRegister)

	c.Status(http.StatusOK)
}

func (u *UserController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		panic("Id not valid format, it must be int")
	}
	res := u.UserService.Delete(uint(id))

	if res {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}

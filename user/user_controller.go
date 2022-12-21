package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService UserService
}

func ProvideUserController(u UserService) UserController {
	return UserController{UserService: u}
}

func (u *UserController) Create(c *gin.Context) {
	var userRegister UserRequest
	err := c.ShouldBind(&userRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	user := RegisterToUser(userRegister)
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

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u *UserController) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		panic("Id not valid format, it must be int")
	}
	var userRegister UserRequest
	err2 := c.ShouldBind(&userRegister)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err2)})
		return
	}

	u.UserService.Update(uint(id), userRegister)

	c.Status(http.StatusOK)
}

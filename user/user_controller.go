package user

import (
	"github.com/gin-gonic/gin"
	"log"
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
	var userRegister UserRegister
	err := c.BindJSON(&userRegister)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
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

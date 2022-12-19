package routes

import (
	"freeDayManager/role"
	"freeDayManager/user"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, roleController *role.RoleController, userController *user.UserController) {
	router.POST("/api/roles", roleController.Create)
	router.POST("/api/register", userController.Register)
}

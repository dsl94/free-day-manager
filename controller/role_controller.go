package controller

import (
	"freeDayManager/dto"
	"freeDayManager/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RoleController struct {
	RoleService service.RoleService
}

func ProvideRoleController(r service.RoleService) RoleController {
	return RoleController{RoleService: r}
}

func (r *RoleController) Create(c *gin.Context) {

	var roleDto dto.RoleDto
	err := c.BindJSON(&roleDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	createdRole := r.RoleService.Save(dto.ToRole(roleDto))

	c.JSON(http.StatusOK, gin.H{"role": dto.ToDto(createdRole)})
}

package controller

import (
	"fmt"
	"freeDayManager/authorizer"
	"freeDayManager/dto"
	"freeDayManager/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FreeDayController struct {
	FreeDayService service.FreeDayService
}

func ProvideFreeDayController(r service.FreeDayService) FreeDayController {
	return FreeDayController{FreeDayService: r}
}

func (f *FreeDayController) Create(c *gin.Context) {
	var freeDayRequestDto dto.FreeDayRequest
	err := c.ShouldBind(&freeDayRequestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	freeDayRequest, err2 := dto.MapToFreeDay(freeDayRequestDto)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err2)})
		return
	}

	username := authorizer.ExtractUsername(c)

	created, err3 := f.FreeDayService.CreateRequest(freeDayRequest, username)
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err3)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"freeDay": created})
}

func (f *FreeDayController) FindAll(c *gin.Context) {
	all := f.FreeDayService.FindAll()
	c.JSON(http.StatusOK, gin.H{"freeDayRequests": all})
}

func (f *FreeDayController) FindForUser(c *gin.Context) {
	username := authorizer.ExtractUsername(c)
	all := f.FreeDayService.FindForUser(username)
	c.JSON(http.StatusOK, gin.H{"freeDayRequests": all})
}

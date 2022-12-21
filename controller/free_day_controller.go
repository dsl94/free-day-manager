package controller

import "freeDayManager/service"

type FreeDayController struct {
	FreeDayService service.FreeDayService
}

func ProvideFreeDayController(r service.FreeDayService) FreeDayController {
	return FreeDayController{FreeDayService: r}
}

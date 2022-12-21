package service

import (
	"freeDayManager/database"
)

type FreeDayService struct {
	FreeDayRepository database.FreeDayRepository
}

func ProvideFreeDayService(r database.FreeDayRepository) FreeDayService {
	return FreeDayService{FreeDayRepository: r}
}

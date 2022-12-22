package service

import (
	"errors"
	"freeDayManager/database"
	"freeDayManager/dto"
	"freeDayManager/entity"
)

type FreeDayService struct {
	FreeDayRepository database.FreeDayRepository
	UserRepository    database.UserRepository
}

func ProvideFreeDayService(r database.FreeDayRepository, ur database.UserRepository) FreeDayService {
	return FreeDayService{FreeDayRepository: r, UserRepository: ur}
}

func (f *FreeDayService) CreateRequest(freeDayRequest entity.FreeDay, username string) (dto.FreeDayResponse, error) {
	dbUser := f.UserRepository.FindByUsername(username)
	if dbUser.ID == 0 {
		return dto.FreeDayResponse{}, errors.New("User does not exist")
	}
	freeDayRequest.User = dbUser

	created := f.FreeDayRepository.Save(freeDayRequest)

	return dto.MapToFreeDayResponse(created), nil
}

func (f *FreeDayService) FindAll() []dto.FreeDayResponse {
	all := f.FreeDayRepository.FindAll()
	var dtos = make([]dto.FreeDayResponse, len(all))

	for i, day := range all {
		dtos[i] = dto.MapToFreeDayResponse(day)
	}

	return dtos
}

func (f *FreeDayService) FindForUser(username string) []dto.FreeDayResponse {
	dbUser := f.UserRepository.FindByUsername(username)
	all := f.FreeDayRepository.FindByUserId(int(dbUser.ID))
	var dtos = make([]dto.FreeDayResponse, len(all))

	for i, day := range all {
		dtos[i] = dto.MapToFreeDayResponse(day)
	}

	return dtos
}

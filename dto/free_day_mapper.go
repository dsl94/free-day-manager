package dto

import (
	"errors"
	"freeDayManager/entity"
	"time"
)

func MapToFreeDay(request FreeDayRequest) (entity.FreeDay, error) {
	startDate, err := time.Parse("2006-01-02", request.StartDate)
	endDate, err1 := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		return entity.FreeDay{}, errors.New(err.Error())
	}
	if err1 != nil {
		return entity.FreeDay{}, errors.New(err1.Error())
	}

	freeDayEntity := entity.FreeDay{
		Reason:    request.Reason,
		Status:    entity.FREE_DAY_REQUESTED,
		StartDate: startDate,
		EndDate:   endDate,
	}

	return freeDayEntity, nil
}

func MapToFreeDayResponse(freeDay entity.FreeDay) FreeDayResponse {
	response := FreeDayResponse{
		ID:        freeDay.ID,
		Reason:    freeDay.Reason,
		Status:    freeDay.Status,
		StartDate: freeDay.StartDate.Format("2006-01-02"),
		EndDate:   freeDay.EndDate.Format("2006-01-02"),
		User:      UserToDto(freeDay.User),
	}

	return response
}

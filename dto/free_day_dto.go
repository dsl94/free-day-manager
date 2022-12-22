package dto

type FreeDayRequest struct {
	Reason    string `json:"reason" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

type FreeDayResponse struct {
	ID        uint    `json:"id"`
	Reason    string  `json:"reason"`
	Status    string  `json:"status"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	User      UserDto `json:"user"`
}

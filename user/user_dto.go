package user

type UserLogin struct {
	Username string
	Password string
}

type UserRequest struct {
	Username string `json:"username" binding:"required_without=End,omitempty,lt|ltfield=End"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required_without=End,omitempty,lt|ltfield=End"`
	FullName string `json:"full_name" binding:"required_without=End,omitempty,lt|ltfield=End"`
	Role     string `json:"role" binding:"required_without=End,omitempty,lt|ltfield=End"`
}

type UserDto struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	FullName string   `json:"full_name"`
	Roles    []string `json:"roles"`
}

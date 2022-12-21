package user

type UserLogin struct {
	Username string
	Password string
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UserDto struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	FullName string   `json:"full_name"`
	Roles    []string `json:"roles"`
}

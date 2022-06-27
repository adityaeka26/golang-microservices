package web

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type GetUserRequest struct {
	Id string `form:"id" validate:"required"`
}

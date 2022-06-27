package web

type GetUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

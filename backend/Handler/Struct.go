package Handlers

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserInfo struct {
	Token string `json:"token"`
}
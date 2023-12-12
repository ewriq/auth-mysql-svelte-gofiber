package Database
type User struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
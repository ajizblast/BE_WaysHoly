package userdto

type UserResponse struct {
	ID       int    `json:"id"`
	FullName     string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
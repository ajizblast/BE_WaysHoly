package authdto

type RegisterResponse struct {
	FullName string `json:"fullName"`
	Token    string `json:"token"`
}

type LoginResponse struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
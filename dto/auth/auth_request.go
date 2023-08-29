package authdto

type AuthRequest struct {
	FullName string `json:"fullName" form:"fullName"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    int    `json:"phone" form:"phone"`
	Image    string `json:"image" form:"image"`
	Address  string `json:"address" form:"address"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

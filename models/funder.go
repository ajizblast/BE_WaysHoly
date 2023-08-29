package models

type Funder struct {
	ID         int                  `json:"id"`
	CreatedAt  string               `json:"donate_at"`
	Total      int                  `json:"total"`
	Status     string               `json:"status"`
	DonationID int                  `json:"donation_id" form:"donation_id"`
	Donation   DonationResponse     `json:"donation" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID     int                  `json:"user_id"`
	User       UsersProfileResponse `json:"user"`
}

type FundersDonationResponse struct {
	ID        int    `json:"id"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	Total     int    `json:"donateAmount"`
	Status    string `json:"status"`
	CreatedAt string `json:"donate_at"`
}

type FunderUser struct {
	ID        int                 `json:"id"`
	CreatedAt string              `json:"donate_at"`
	Total     int                 `json:"total"`
	UserID    int                 `json:"user_id"`
	User      UsersFunderResponse `json:"user"`
}

func (FunderUser) TableName() string {
	return "funders"
}

func (FundersDonationResponse) TableName() string {
	return "funders"
}
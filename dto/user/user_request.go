package userdto

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" type:"varchar(255)"`
	Phone    int    `json:"phone" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
}
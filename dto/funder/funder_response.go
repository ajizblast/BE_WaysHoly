package funderdto

import (
	"holyways/models"
	"time"
)

type FunderUserResponse struct {
	ID        int                        `json:"id"`
	UserID    int                        `json:"user_id"`
	User      models.UsersFunderResponse `json:"user"`
	Total     int                        `json:"total"`
	CreatedAt time.Time                  `json:"donate_at"`
}
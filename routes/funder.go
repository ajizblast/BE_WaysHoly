package routes

import (
	"holyways/handlers"
	"holyways/pkg/mysql"
	"holyways/repository"
	"holyways/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func FunderRoutes(e *echo.Group) {
	FunderRepository := repository.RepositoryFunder(mysql.DB)

	h := handlers.HandlerFunder(FunderRepository)

	e.GET("/funders", h.FindFunder)
	e.GET("/funder", middleware.Auth(h.GetFunder))
	e.GET("/funder-by-login", middleware.Auth(h.FindFunderByLogin))
	e.GET("/funder-by-donation-and-status-succes/:id", h.FindFunderByDonationIDAndStatusSucces)
	e.GET("/funder-by-donation-and-status-pending/:id", h.FindFunderByDonationIDAndStatusPending)
	e.POST("/funder", middleware.Auth(h.CreateFunder))
	e.POST("/notification", h.Notification)
}

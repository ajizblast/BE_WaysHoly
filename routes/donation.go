package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	"holyways/pkg/mysql"
	"holyways/repository"

	"github.com/labstack/echo/v4"
)

func DonationRoute(e *echo.Group) {
	donationRepository := repository.RepositoryDonation(mysql.DB)
	h := handlers.HandlerDonation(donationRepository)

	e.GET("/donations", h.FindDonation)
	e.GET("/donation/:id", h.GetDonation)
	e.GET("/donation-by-user", middleware.Auth(h.GetDonationByUserID))
	e.POST("/donation", middleware.Auth(middleware.UploadFile(h.CreateDonation)))
	e.PATCH("/donation/:id", middleware.Auth(middleware.UploadFile(h.Updatedonation)))
	e.DELETE("/donation/:id", middleware.Auth(h.DeleteDonation))
}
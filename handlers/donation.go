package handlers

import (
	"context"
	"fmt"
	donationdto "holyways/dto/donation"
	dto "holyways/dto/result"
	"holyways/models"
	"holyways/repository"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerDonation struct {
	DonationRepository repository.DonationRepository
}

func HandlerDonation(DonationRepository repository.DonationRepository) *handlerDonation {
	return &handlerDonation{DonationRepository}
}

func (h *handlerDonation) FindDonation(c echo.Context) error {
	donations, err := h.DonationRepository.FindDonation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: donations})
}

func (h *handlerDonation) GetDonation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	donation, err := h.DonationRepository.GetDonation(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: donation})
}

func (h *handlerDonation) CreateDonation(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	goal, _ := strconv.Atoi(c.FormValue("goal"))

	request := donationdto.DonationRequest{
		Title:       c.FormValue("title"),
		Goal:        goal,
		Description: c.FormValue("description"),
		Thumbnail:   dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "holyways"})
	if err != nil {
		fmt.Println(err.Error())
	}
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	donation := models.Donation{
		Title:       request.Title,
		Goal:        request.Goal,
		Description: request.Description,
		Thumbnail:   resp.SecureURL,
		UserID:      int(userId),
	}

	donation, err = h.DonationRepository.CreateDonation(donation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	donation, _ = h.DonationRepository.GetDonation(donation.ID)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDonation(donation)})
}

func (h *handlerDonation) GetDonationByUserID(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	donation, err := h.DonationRepository.GetDonationByUserID(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: donation})
}

func (h *handlerDonation) Updatedonation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	dataFile := c.Get("dataFile").(string)

	goal, _ := strconv.Atoi(c.FormValue("goal"))

	request := donationdto.DonationRequest{
		Title:       c.FormValue("title"),
		Goal:        goal,
		Description: c.FormValue("description"),
		Thumbnail:   dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "holyways"})
	if err != nil {
		fmt.Println(err.Error())
	}

	donation, _ := h.DonationRepository.GetDonation(id)

	if request.Title != "" {
		donation.Title = request.Title
	}

	if request.Goal != 0 {
		donation.Goal = request.Goal
	}

	if request.Description != "" {
		donation.Description = request.Description
	}

	if request.Thumbnail != "" {
		donation.Thumbnail = resp.SecureURL
	}

	data, err := h.DonationRepository.UpdateDonation(donation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDonation(data)})
}

func (h *handlerDonation) DeleteDonation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	donation, err := h.DonationRepository.GetDonation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.DonationRepository.DeleteDonation(donation, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: ConvertResponseDeleteDonation(data)})
}

func ConvertResponseDeleteDonation(d models.Donation) donationdto.DonationDeleteResponse {
	return donationdto.DonationDeleteResponse{
		ID: d.ID,
	}
}

func convertResponseDonation(u models.Donation) models.DonationResponse {
	return models.DonationResponse{
		ID:          u.ID,
		Title:       u.Title,
		User:        u.User,
		Goal:        u.Goal,
		Description: u.Description,
		Thumbnail:   u.Thumbnail,
	}
}
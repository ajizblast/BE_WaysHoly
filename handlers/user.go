package handlers

import (
	dto "holyways/dto/result"
	userdto "holyways/dto/user"
	"holyways/models"
	"holyways/repository"
	"net/http"
	"strconv"
	"holyways/pkg/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// struct save connetion
type handler struct {
	UserRepository repository.UserRepository
}

// function connection
func HandlerUser(UserRepository repository.UserRepository) *handler {
	return &handler{UserRepository}
}

// find all users
func (h *handler) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: users})
}

func (h *handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponse(user)})
}

func (h *handler) GetUserIDByLogin(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}


func (h *handler) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	phone, _ := strconv.Atoi(c.FormValue("phone"))

	request := userdto.UpdateUserRequest{
		FullName: c.FormValue("fullName"),
		Email:    c.FormValue("email"),
		Image:    c.FormValue("image"),
		Phone:    phone,
		Address:  c.FormValue("address"),
	}

	// bcrypt pasword
	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.FullName != "" {
		user.FullName = request.FullName
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = password
	}

	if request.Image != "" {
		user.Image = request.Image
	}

	if request.Phone != 0 {
		user.Phone = request.Phone
	}

	if request.Address != "" {
		user.Address = request.Address
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}


// DELETE USER
func (h *handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

// convert response data
func convertResponse(user models.User) userdto.UserResponse {
	return userdto.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
	}
}
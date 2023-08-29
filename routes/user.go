package routes

import (
	"holyways/handlers"
	"holyways/pkg/mysql"
	"holyways/repository"
	"holyways/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repository.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	e.GET("/user-by-login",middleware.Auth(h.GetUserIDByLogin))
	e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.DeleteUser)
}
package main

import (
	"fmt"
	"holyways/database"
	"holyways/pkg/mysql"
	"holyways/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	// CORS agar bisa akses backend
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},                                                 // mengijinkan akses semuanya
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},        // mengijinkan method apa aja yg bisa digunakan
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"}, // mengijinkan headers apa aja yg bisa digunakan
	}))

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	var PORT = os.Getenv("PORT")

	fmt.Println("server running localhost:" + PORT)
	e.Logger.Fatal(e.Start(":" + PORT)) // delete localhost

	// fmt.Println("Server running on localhost:5000")
	// e.Logger.Fatal(e.Start("localhost:5000"))

	// var PORT = os.Getenv("PORT")

	// // fmt.Println("server running localhost:" + PORT)
	// e.Logger.Fatal(e.Start(":" + PORT)) // delete localhost
}

package database

import (
	"holyways/models"
	"holyways/pkg/mysql"
	"fmt"
)

// automatic migrate
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Donation{},
		&models.Funder{},
	)

	if err != nil {
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
package repository

import (
	"holyways/models"
	"gorm.io/gorm"
)
	


type DonationRepository interface{
	FindDonation() ([]models.DonationResponse, error)
	GetDonation(ID int) (models.Donation, error)
	CreateDonation(donation models.Donation) (models.Donation, error)
	GetDonationByUserID(userId int) ([]models.Donation, error)
	UpdateDonation(donation models.Donation) (models.Donation, error)
	DeleteDonation(donation models.Donation, ID int) (models.Donation, error)
}

func RepositoryDonation(db *gorm.DB) *repositories {
	return &repositories{db}
}
func (r *repositories) FindDonation() ([]models.DonationResponse, error) {
	var donations []models.DonationResponse
	err := r.db.Preload("User").Find(&donations).Error

	return donations, err
}
func (r *repositories) GetDonation(ID int) (models.Donation, error) {
	var donation models.Donation
	err := r.db.Preload("User").First(&donation, ID).Error

	return donation, err
}
func (r *repositories) CreateDonation(donation models.Donation) (models.Donation, error) {
	err := r.db.Create(&donation).Error

	return donation, err
}

func (r *repositories) GetDonationByUserID(userId int) ([]models.Donation, error) {
	var donations []models.Donation
	err := r.db.Where("user_id=?", userId).Preload("User").Find(&donations).Error
	return donations, err
}

func (r *repositories) UpdateDonation(donation models.Donation) (models.Donation, error) {
	err := r.db.Save(&donation).Error // Using Save method ORM
  
	return donation, err
}
func (r *repositories) DeleteDonation(donation models.Donation , ID int) (models.Donation, error) {
	err := r.db.Delete(&donation).Error // Using Delete method ORM
	
	return donation, err
}
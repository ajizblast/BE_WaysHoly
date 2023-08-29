package repository

import (
	"holyways/models"

	"gorm.io/gorm"
)

type FunderRepository interface {
	FindFunder() ([]models.Funder, error)
	FindFunderByLogin(userId int) ([]models.Funder, error)
	FindFunderByStatusSucces(userId int) ([]models.Funder, error)
	FindFunderByDonationIDAndStatusSucces(donationId int) ([]models.Funder, error)
	FindFunderByDonationIDAndStatusPending(donationId int) ([]models.Funder, error)
	GetFunder(ID int) (models.Funder, error)
	GetFunderID(funderId int) (models.Funder, error)
	GetFunderByDonation(ID int) ([]models.Funder, error)
	CreateFunder(transaction models.Funder) (models.Funder, error)
	UpdateFunder(status string, orderId int) (models.Funder, error)
}

func RepositoryFunder(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindFunder() ([]models.Funder, error) {
	var funders []models.Funder
	err := r.db.Preload("Donation.User").Preload("User").Find(&funders).Error

	return funders, err
}

func (r *repositories) FindFunderByLogin(userId int) ([]models.Funder, error) {
	var funders []models.Funder
	err := r.db.Where("user_id=?", userId).Preload("Donation").Preload("User").Find(&funders).Error

	return funders, err
}

func (r *repositories) FindFunderByDonationIDAndStatusSucces(donationId int) ([]models.Funder, error) {
	var funders []models.Funder
	err := r.db.Where("donation_id=?", donationId).Where("status=?", "success").Preload("Donation").Preload("User").Find(&funders).Error

	return funders, err
}

func (r *repositories) FindFunderByDonationIDAndStatusPending(donationId int) ([]models.Funder, error) {
	var funders []models.Funder
	err := r.db.Where("donation_id=?", donationId).Where("status=?", "pending").Preload("Donation").Preload("User").Find(&funders).Error

	return funders, err
}


func (r *repositories) FindFunderByStatusSucces(userId int) ([]models.Funder, error) {
	var funders []models.Funder
	err := r.db.Where("user_id=?", userId).Where("status=?", "success").Preload("Donation").Preload("User").Find(&funders).Error

	return funders, err
}

func (r *repositories) GetFunder(ID int) (models.Funder, error) {
	var funder models.Funder
	err := r.db.Preload("Donation").Preload("User").First(&funder, ID).Error

	return funder, err
}

func (r *repositories) GetFunderID(funderId int) (models.Funder, error) {
	var funder models.Funder
	err := r.db.Preload("Donation").Preload("User").First(&funder, funderId).Error

	return funder, err
}

func (r *repositories) GetFunderByDonation(ID int) ([]models.Funder, error) {
	var funders []models.Funder
	err := r.db.Where("donation_id", ID).Find(&funders).Error

	return funders, err
}

func (r *repositories) CreateFunder(funder models.Funder) (models.Funder, error) {

	err := r.db.Preload("Donation").Preload("User").Create(&funder).Error

	return funder, err
}

func (r *repositories) UpdateFunder(status string, orderId int) (models.Funder, error) {
	var funder models.Funder

	r.db.Preload("Donation").Preload("User").First(&funder, orderId)
	if status != funder.Status && status == "success" {
		var donation models.Donation
		r.db.First(&donation, funder.Donation.ID)
		donation.CurrentGoal = donation.CurrentGoal + funder.Total
		donation.Goal = donation.Goal + 1
		r.db.Save(&donation)
	}

	funder.Status = status
	err := r.db.Save(&funder).Error

	return funder, err
}

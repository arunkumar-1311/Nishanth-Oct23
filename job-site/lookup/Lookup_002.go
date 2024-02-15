package lookup

import (
	"job-post/models"
	"job-post/service/helper"

	"gorm.io/gorm"
)

func (Empty) Lookup_002(db *gorm.DB) error {
	roles := []models.Roles{
		{Role_ID: "AD1", Role: "Admin"},
		{Role_ID: "USER1", Role: "User"},
	}

	jobType := []models.JobType{
		{JobTypeID: "JT-CON", JobType: "Contract"},
		{JobTypeID: "JT-FT", JobType: "Full-Time"},
		{JobTypeID: "JT-FL", JobType: "FreeLance"},
		{JobTypeID: "JT-TC", JobType: "Telecom"},
	}

	country := []models.Country{
		{CountryID: helper.UniqueID(), Country: "India"},
		{CountryID: helper.UniqueID(), Country: "Australia"},
		{CountryID: helper.UniqueID(), Country: "United Kingdom"},
		{CountryID: helper.UniqueID(), Country: "United Status"},
		{CountryID: helper.UniqueID(), Country: "Canada"},
	}

	if err := db.Create(roles).Error; err != nil {
		return err
	}

	if err := db.Create(jobType).Error; err != nil {
		return err
	}

	if err := db.Create(country).Error; err != nil {
		return err
	}
	return nil
}

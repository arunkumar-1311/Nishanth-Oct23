package repository

import "job-post/models"

type Country interface {
	ReadCountries(*[]models.Country) error
}

// Helps to read all the countries and its ID
func (d *GORM_Connection) ReadCountries(dest *[]models.Country) error {
	if err := d.DB.Find(&dest).Error; err != nil {
		return err
	}
	return nil
}

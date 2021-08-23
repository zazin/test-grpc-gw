package models

import (
	"gorm.io/gorm"
)

type Country struct {
	ID   int
	Code string
	Name string
}

func (Country) TableName() string {
	return "country"
}

func ListCountry(db *gorm.DB, Country *[]Country) (err error) {
	err = db.Find(Country).Error
	if err != nil {
		return err
	}
	return nil
}

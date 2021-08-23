package country

import (
	"gorm.io/gorm"
	"test-grpc-gw/internal/location/database"
	"test-grpc-gw/internal/location/models"
)

type CountryRepository struct {
	Db *gorm.DB
}

func New() *CountryRepository {
	db := database.InitDb()
	err := db.AutoMigrate(&models.Country{})
	if err != nil {
		return nil
	}
	return &CountryRepository{Db: db}
}

//get countries
func (repository *CountryRepository) GetCountries() []models.Country {
	var countries []models.Country
	err := models.ListCountry(repository.Db, &countries)
	if err != nil {
		return nil
	}
	return countries
}

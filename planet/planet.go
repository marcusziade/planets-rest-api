package planet

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Service struct {
	Database *gorm.DB
}

type Planet struct {
	Name        string
	Distance    int
	Temperature int
	CreatedAt   time.Time
}

type PlanetService interface {
	GetAllPlanets() ([]Planet, error)
	GetPlanet(ID uint) (Planet, error)
	AddPlanet(planet Planet) (Planet, error)
	UpdatePlanet(ID uint, updatedPlanet Planet) (Planet, error)
	DeletePlanet(ID uint) error
}

// Returns a new Planet service
func NewService(database *gorm.DB) *Service {
	return &Service{
		Database: database,
	}
}

// Returns all planets in the database
func (s *Service) GetAllPlanets() ([]Planet, error) {
	var planets []Planet
	if result := s.Database.Find(&planets); result.Error != nil {
		return planets, result.Error
	}

	return planets, nil
}

// Returns a planet for ID
func (s *Service) GetPlanet(ID uint) (Planet, error) {
	var planet Planet
	if result := s.Database.Find(&planet); result.Error != nil {
		return planet, result.Error
	}

	return planet, nil
}

// Adds a new planet to the database
func (s *Service) AddPlanet(planet Planet) (Planet, error) {
	if result := s.Database.Save(&planet); result.Error != nil {
		return Planet{}, result.Error
	}

	return planet, nil
}

// Update an existing planet in the database
func (s *Service) UpdatePlanet(ID uint, updatedPlanet Planet) (Planet, error) {
	planet, err := s.GetPlanet(ID)
	if err != nil {
		return Planet{}, err
	}

	if result := s.Database.Model(&planet).Updates(updatedPlanet); result.Error != nil {
		return Planet{}, result.Error
	}

	return planet, nil
}

// Delete a planet from the database
func (s *Service) DeletePlanet(ID uint) error {
	if result := s.Database.Delete(&Planet{}, ID); result.Error != nil {
		return result.Error
	}

	return nil
}

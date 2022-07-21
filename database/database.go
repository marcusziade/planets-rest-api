package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type ConnectionParameters struct {
	Username string
	Password string
	Host     string
	Database string
	Port     string
}

// Returns a pointer to a new database instance
func NewDatabase(parameters ConnectionParameters) (*gorm.DB, error) {
	// dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", parameters.Host, parameters.Port, parameters.Username, parameters.Database, parameters.Password)
	fmt.Println(connectionString)

	database, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return database, err
	}

	if err := database.DB().Ping(); err != nil {
		return database, err
	}

	return database, nil
}

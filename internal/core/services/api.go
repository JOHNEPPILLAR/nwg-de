// Package services - Establishes communication between the core and the outside world
package services

import (
	"errors"
	"os"
	"time"

	"github.com/JOHNEPPILLAR/nwg-de/internal/core/domain"
	"github.com/JOHNEPPILLAR/nwg-de/internal/core/ports"
	"github.com/JOHNEPPILLAR/nwg-de/internal/utility"
)

// APIService -
type APIService struct {
	logger utility.Logger
	repo   ports.APIRepository
}

// NewAPIService -
func NewAPIService(logger utility.Logger, repo ports.APIRepository) *APIService {
	return &APIService{
		logger: logger,
		repo:   repo,
	}
}

// StartAPIServer -
func (a *APIService) StartAPIServer() error {
	err := a.StartAPIServer()
	if err != nil {
		a.logger.Fatal("Unable to start api server")
		os.Exit(1)
	}

	return nil
}

//
// Connect API Routes
//

// HealthCheck endpoint -
func (a *APIService) HealthCheck() error {
	return a.repo.HealthCheck()
}

// AddVehicle endpoint -
func (a *APIService) AddVehicle(vehicle *domain.Vehicle) error {
	return a.repo.AddVehicle(vehicle)
}

// FindVehicle endpoint -
func (a *APIService) FindVehicle(licenseNumber string) ([]*domain.Vehicle, error) {
	return a.repo.FindVehicle(licenseNumber)
}

// GetAllVehicles endpoint -
func (a *APIService) GetAllVehicles() ([]*domain.Vehicle, error) {
	return a.repo.GetAllVehicles()
}

// GetAvailableVehicles endpoint -
func (a *APIService) GetAvailableVehicles(startDate time.Time, endDate time.Time) ([]*domain.Vehicle, error) {
	return a.repo.GetAvailableVehicles(startDate, endDate)
}

// AddBooking endpoint -
func (a *APIService) AddBooking(booking *domain.Booking) error {

	// Get vehicle
	licenseNumber := booking.LicenseNumber
	vehicle, err := a.repo.FindVehicle(licenseNumber)
	if err != nil {
		a.logger.Error(err.Error())
		return err
	}

	if len(vehicle) == 0 {
		err = errors.New("No vehicle found")
	}

	// Check vehicle availability
	if vehicle[0].Reserved {
		err := errors.New("[" + vehicle[0].LicenseNumber + "] Already booked")
		a.logger.Info(err.Error())
		return err
	}

	// Work out booking duration in days
	difference := booking.EndDate.Sub(booking.StartDate)
	duration := difference.Hours() / 24

	// Work out total booking cost
	booking.Cost = vehicle[0].CostPerDay * float64(duration)

	// Update vehicle's booked status
	err = a.repo.BookVehicle(vehicle[0])
	if err != nil {
		a.logger.Error(err.Error())
		return err
	}

	// Save booking
	return a.repo.AddBooking(booking)
}

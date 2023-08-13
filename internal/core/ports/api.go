// Package ports - host application contracts/interfaces
package ports

import (
	"time"

	"github.com/JOHNEPPILLAR/nwg-de/internal/core/domain"
)

//
// Driving Adaptors
//

// APIService -
type APIService interface {
	StartAPIServer() error
	HealthCheck() error
	AddVehicle(*domain.Vehicle) error
	FindVehicle(string) ([]*domain.Vehicle, error)
	GetAllVehicles() ([]*domain.Vehicle, error)
	GetAvailableVehicles() ([]*domain.Vehicle, error)

	AddBooking(*domain.Booking) error
}

//
// Driven Adaptors
//

// APIRepository -
type APIRepository interface {
	HealthCheck() error
	AddVehicle(*domain.Vehicle) error
	BookVehicle(*domain.Vehicle) error
	FindVehicle(string) ([]*domain.Vehicle, error)
	GetAllVehicles() ([]*domain.Vehicle, error)
	GetAvailableVehicles(startDate time.Time, endDate time.Time) ([]*domain.Vehicle, error)

	AddBooking(*domain.Booking) error
}

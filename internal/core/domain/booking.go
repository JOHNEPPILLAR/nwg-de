// Package domain - Data models
package domain

import "time"

// Booking -
type Booking struct {
	LicenseNumber string
	StartDate     time.Time
	EndDate       time.Time
	Cost          float64
}

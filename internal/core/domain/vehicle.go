// Package domain - Data models
package domain

// Vehicle -
type Vehicle struct {
	Make          string
	Model         string
	SunRoof       bool
	TowBar        bool
	LicenseNumber string
	CostPerDay    float64
	Reserved      bool
}

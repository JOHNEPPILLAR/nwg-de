// Package repository - Mongo data store
package repository

import (
	"time"

	"github.com/JOHNEPPILLAR/nwg-de/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAvailableVehicles -
func (m *MongoRepository) GetAvailableVehicles(startDate time.Time, endDate time.Time) ([]*domain.Vehicle, error) {

	// Set DB & Collection
	db := m.client.Database("nwg-de")
	dbCollection := db.Collection("bookings")

	// Set options
	opts := options.Find()
	filter := bson.M{
		"startdate": bson.M{"$gte": startDate.UTC()},
		"enddate":   bson.M{"$lt": endDate.AddDate(0, 0, 1).UTC()},
	}

	// Get data from DB
	cursor, err := dbCollection.Find(m.ctx, filter, opts)
	defer cursor.Close(m.ctx)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}

	// Convert into strut
	var bookedVehicle []*domain.Booking
	if err = cursor.All(m.ctx, &bookedVehicle); err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}

	// Filter out booked vehicles
	results := []*domain.Vehicle{}
	allVehicles, err := m.GetAllVehicles()
	for _, vehicle := range allVehicles {
		for _, booking := range bookedVehicle {
			if booking.LicenseNumber != vehicle.LicenseNumber {
				results = append(results, vehicle)
			}
		}
	}

	// Return [] if nothing is available
	if len(results) == 0 {
		return []*domain.Vehicle{}, nil
	}

	return results, nil
}

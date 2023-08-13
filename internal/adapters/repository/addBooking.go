// Package repository - Mongo data store
package repository

import (
	"errors"

	"github.com/JOHNEPPILLAR/nwg-de/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddBooking -
func (m *MongoRepository) AddBooking(booking *domain.Booking) error {

	// Set DB & collection
	db := m.client.Database("nwg-de")
	dbCollection := db.Collection("bookings")

	// Get data from DB
	opts := options.Count()
	filter := bson.D{{Key: "licensenumber", Value: booking.LicenseNumber}}
	count, err := dbCollection.CountDocuments(m.ctx, filter, opts)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}

	if count == 1 {
		err := errors.New("[" + booking.LicenseNumber + "] Booking already exists")
		m.logger.Debug(err.Error())
		return err
	} else {
		// Save data
		_, err := dbCollection.InsertOne(m.ctx, booking)
		if err != nil {
			m.logger.Error("[" + booking.LicenseNumber + "] Failed to save")
			return err
		}
	}

	m.logger.Info("[" + booking.LicenseNumber + "] Booking added")

	return nil
}

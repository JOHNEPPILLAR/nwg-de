// Package repository - Mongo data store
package repository

import (
	"errors"

	"github.com/JOHNEPPILLAR/nwg-de/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddVehicle -
func (m *MongoRepository) AddVehicle(vehicle *domain.Vehicle) error {

	// Set DB & collection
	db := m.client.Database("nwg-de")
	dbCollection := db.Collection("vehicles")

	// Get data from DB
	opts := options.Count()
	filter := bson.D{{Key: "licensenumber", Value: vehicle.LicenseNumber}}
	count, err := dbCollection.CountDocuments(m.ctx, filter, opts)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}

	if count == 1 {
		err := errors.New("[" + vehicle.LicenseNumber + "] Already exists")
		m.logger.Debug(err.Error())
		return err
	} else {
		// Save data
		_, err := dbCollection.InsertOne(m.ctx, vehicle)
		if err != nil {
			m.logger.Error("[" + vehicle.LicenseNumber + "] Failed to save")
			return err
		}
	}

	m.logger.Info("[" + vehicle.LicenseNumber + "] Added")

	return nil
}

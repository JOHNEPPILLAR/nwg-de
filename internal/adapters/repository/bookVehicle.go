// Package repository - Mongo data store
package repository

import (
	"github.com/JOHNEPPILLAR/nwg-de/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// BookVehicle -
func (m *MongoRepository) BookVehicle(vehicle *domain.Vehicle) error {

	// Set DB & collection
	db := m.client.Database("nwg-de")
	dbCollection := db.Collection("vehicles")

	// Get data from DB
	filter := bson.D{{Key: "licensenumber", Value: vehicle.LicenseNumber}}
	update := bson.M{"$set": bson.M{"reserved": !vehicle.Reserved}}

	// Save data
	result, err := dbCollection.UpdateOne(m.ctx, filter, update)
	if err != nil {
		m.logger.Error("[" + vehicle.LicenseNumber + "] Failed to update booking status")
		return err
	}

	if result.ModifiedCount > 0 {
		m.logger.Info("[" + vehicle.LicenseNumber + "] Vehicle booking status updated")
	}

	return nil
}

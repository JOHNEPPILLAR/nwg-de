// Package repository - Mongo data store
package repository

import (
	"github.com/JOHNEPPILLAR/nwg-de/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindVehicle -
func (m *MongoRepository) FindVehicle(licenseNumber string) ([]*domain.Vehicle, error) {

	// Set DB & Collection
	db := m.client.Database("nwg-de")
	dbCollection := db.Collection("vehicles")

	// Set options
	opts := options.Find()
	filter := bson.D{{Key: "licensenumber", Value: licenseNumber}}

	// Get data from DB
	cursor, err := dbCollection.Find(m.ctx, filter, opts)
	defer cursor.Close(m.ctx)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}

	// Convert into strut
	var results []*domain.Vehicle
	if err = cursor.All(m.ctx, &results); err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}

	// Return results, [] if nothing in DB
	if len(results) == 0 {
		return []*domain.Vehicle{}, nil
	}

	return results, nil
}

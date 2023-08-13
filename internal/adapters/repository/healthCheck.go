// Package repository - Mongo data store
package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// HealthCheck -
func (m *MongoRepository) HealthCheck() error {

	// Set DB
	db := m.client.Database("nwg-de")

	// Ping DB
	if err := db.Client().Ping(context.Background(), readpref.Primary()); err != nil {
		return err
	}

	return nil
}

// Package repository - Mongo data store
package repository

import (
	"context"
	"os"

	"github.com/JOHNEPPILLAR/nwg-de/internal/utility"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoRepository -
type MongoRepository struct {
	client *mongo.Client
	ctx    context.Context
	logger utility.Logger
}

// NewMongoAdaptor -
func NewMongoAdaptor(ctx context.Context, logger utility.Logger) *MongoRepository {

	client, err := ConnectDB(ctx, logger)
	if err != nil {
		os.Exit(1)
	}

	return &MongoRepository{
		client: client,
		ctx:    ctx,
		logger: logger,
	}
}

// ConnectDB - Connect to mongo
func ConnectDB(ctx context.Context, logger utility.Logger) (*mongo.Client, error) {

	// Get connection info
	databaseKey := "DATABASE"

	dbURL, err := utility.GetVaultSecret(databaseKey)
	if err != nil {
		return nil, err
	}

	logger.Debug("Connecting to Mongo...")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// Check Mongo is online
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Debug("Connected to Mongo")

	return client, nil
}

// Disconnect -
func (m *MongoRepository) Disconnect() error {
	m.logger.Warn("Disconnecting database...")
	return m.client.Disconnect(m.ctx)
}

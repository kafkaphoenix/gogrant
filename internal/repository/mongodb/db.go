package mongodb

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoDB struct {
	logger   *slog.Logger
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDB(logger *slog.Logger) *MongoDB {
	return &MongoDB{
		logger: logger.With("component", "mongodb"),
	}
}

func (db *MongoDB) Connect(ctx context.Context, uri string) error {
	db.logger.Debug("Connecting to database...", "uri", uri)

	if uri == "" {
		return ErrEmptyURI
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return &MongoDBError{
			Message: "failed to connect to MongoDB",
			Err:     err,
		}
	}

	db.logger.Debug("Pinging to database...", "uri", uri)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return &MongoDBError{
			Message: "failed to ping MongoDB",
			Err:     err,
		}
	}

	db.client = client
	db.database = client.Database("gogrant")
	db.logger.Debug("Connected to MongoDB", "uri", uri)

	return nil
}

func (db *MongoDB) Disconnect(ctx context.Context) {
	if db.client == nil {
		return // No client to disconnect
	}

	if err := db.client.Disconnect(ctx); err != nil {
		db.logger.Error("failed to disconnect from MongoDB", "error", err)
	}

	db.logger.Info("disconnected from MongoDB")
}

func (db *MongoDB) Collection(name string) *mongo.Collection {
	if db.database == nil {
		return nil // No database available
	}

	return db.database.Collection(name)
}

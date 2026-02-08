package mongodb

import (
	"context"
	"os"
	"quentinha_golang/src/configuration/logger"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGODBURL = "MONGODBURL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error)	 {
	mongodb_uri := os.Getenv(MONGODBURL)
	mongodb_database := os.Getenv(MONGODB_USER_DB)
	client, err := mongo.Connect( 
		options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	logger.Info("Connected to MongoDB!")

	return client.Database(mongodb_database), nil

}

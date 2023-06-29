package mongodb

import (
	"context"
	"os"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL      = "MONGODB_URL"
	MONGODB_DATABASE = "MONGODB_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbUrL := os.Getenv(MONGODB_URL)
	mongodbDatabase := os.Getenv(MONGODB_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUrL))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	logger.Info("conexão realizada com sucesso!")

	return client.Database(mongodbDatabase), nil
}

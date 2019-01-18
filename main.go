package main

import (
	"context"
	"fmt"

	"github.com/Noah-Huppert/exercise-tracker/config"

	"github.com/Noah-Huppert/golog"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	ctx := context.Background()

	// {{{1 Create logger
	logger := golog.NewStdLogger("exercise-tracker")

	// {{{1 Load configuration
	cfg, err := config.NewFromEnv()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	logger.Debugf("loaded configuration")

	// {{{1 Connect to MongoDB
	// connect
	mongoClient, err := mongo.NewClient(fmt.Sprintf("mongodb://%s:%d",
		cfg.MongoDBHost, cfg.MongoDBPort))

	// ping
	err = mongoClient.Ping(ctx, nil)

	if err != nil {
		logger.Fatalf("error testing MongoDB connection: %s", err.Error())
	}

	logger.Debugf("connected to MongoDB")
}

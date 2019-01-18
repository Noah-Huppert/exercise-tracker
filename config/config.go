package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds configuration
type Config struct {
	// MongoHost is host at which Mongo can be accessed
	MongoHost string

	// MongoPort is the port at which Mongo can be accessed
	MongoPort int64
}

// NewFromEnv creates a new Config from environment variables
func NewFromEnv() (*Config, error) {
	c := Config{}

	// {{{1 Load values from env vars
	// MongoHost
	c.MongoHost = os.Getenv("MONGO_HOST")

	if len(c.MongoHost) == 0 {
		return nil, fmt.Errorf("MONGO_HOST env var cannot be empty")
	}

	// MongoPort
	mongodbPortStr := os.Getenv("MONGO_PORT")

	if len(mongodbPortStr) == 0 {
		return nil, fmt.Errorf("MONGO_PORT env var cannot be empty")
	}

	mongodbPort, err := strconv.ParseInt(mongodbPortStr, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("error converting MONGO_PORT to int64: %s",
			err.Error())
	}

	c.MongoPort = mongodbPort

	// {{{1 Done
	return &c, nil
}

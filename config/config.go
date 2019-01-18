package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds configuration
type Config struct {
	// MongoDBHost is host at which MongoDB can be accessed
	MongoDBHost string

	// MongoDBPort is the port at which MongoDB can be accessed
	MongoDBPort int64
}

// NewFromEnv creates a new Config from environment variables
func NewFromEnv() (*Config, error) {
	c := Config{}

	// {{{1 Load values from env vars
	// MongoDBHost
	c.MongoDBHost = os.Getenv("MONGODB_HOST")

	if len(c.MongoDBHost) == 0 {
		return nil, fmt.Errorf("MONGODB_HOST env var cannot be empty")
	}

	// MongoDBPort
	mongodbPortStr := os.Getenv("MONGODB_PORT")

	if len(mongodbPortStr) == 0 {
		return nil, fmt.Errorf("MONGODB_PORT env var cannot be empty")
	}

	mongodbPort, err := strconv.ParseInt(mongodbPortStr, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("error converting MONGODB_PORT to int64: %s",
			err.Error())
	}

	c.MongoDBPort = mongodbPort

	// {{{1 Done
	return &c, nil
}

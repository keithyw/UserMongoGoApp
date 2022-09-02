package config

import (
	"context"
	"os"
	"strconv"
	"time"
)

type Config struct {
	MongoDBUri string
	DB string
	Collection string
	Port string
	Context context.Context
}

func NewConfig() (*Config, error) {
	timeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT_BASE"))
	if err != nil {
		return nil, err
	}
	timeoutDuration := time.Duration(timeout * int(time.Second))
	ctx, _ := context.WithTimeout(context.Background(), timeoutDuration)
	return &Config{
		MongoDBUri: os.Getenv("MONGODB_URI"),
		DB: os.Getenv("DB"),
		Collection: os.Getenv("COLLECTION"),
		Port: os.Getenv("PORT"),
		Context: ctx,
	}, nil
}
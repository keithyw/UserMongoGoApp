package database

import (
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"messageapp/config"
)

type UserMongoDB struct {
	DB *mongo.Database
	config *config.Config
	client *mongo.Client
}

func NewMongoDB(config *config.Config) UserMongoDB {
	return UserMongoDB{nil, config, nil}
}

func (um *UserMongoDB) Init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(um.config.MongoDBUri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(um.config.Context)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(um.config.Context, nil)
	if err != nil {
		log.Fatal(err)
	}
	um.client = client
	um.DB = client.Database(um.config.DB)
}
package main

import (
	"context"
	"gitlab.com/hack-city-net/city-net/backend"
	"gitlab.com/hack-city-net/city-net/backend/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Could not create mongodb client: %v", err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatalf("Could not connect to mongodb: %v", err)
	}

	repo, err := store.NewMongodbRepository(client, "city-net", "store")

	proxy := backend.NewLocalReverseProxy(repo)
	proxy.Proxy("0.0.0.0:80")
}

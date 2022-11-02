package main

import (
	"context"
	"gitlab.com/hack-city-net/city-net/backend/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"os"
)

var stores = []store.Store{
	{
		Subdomain:   "zinser.city-offenburg.de",
		WebsiteType: store.WebsiteTypeRedirect,
		Content:     nil,
		Redirect:    "www.mode-zinser.de/home",
	},
	{
		Subdomain:   "mueller.city-offenburg.de",
		WebsiteType: store.WebsiteTypeRedirect,
		Content:     nil,
		Redirect:    "https://www.mueller.de",
	},
	{
		Subdomain:   "test.city-offenburg.de",
		WebsiteType: store.WebsiteTypeHtml,
		Content:     loadContent("backend/test/test1.html"),
		Redirect:    "",
	},
	{
		Subdomain:   "maigarden.city-offenburg.de",
		WebsiteType: store.WebsiteTypePdf,
		Content:     loadContent("backend/test/test2.pdf"),
		Redirect:    "",
	},
}

func loadContent(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("%v", err)
	}
	c, err := io.ReadAll(f)
	if err != nil {
		log.Printf("%v", err)
	}
	_ = f.Close()
	return c
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://188.34.194.85:27017"))
	if err != nil {
		log.Fatalf("Could not create mongodb client: %v", err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatalf("Could not connect to mongodb: %v", err)
	}

	repo, err := store.NewMongodbRepository(client, "city-net", "store")
	for i := 0; i < len(stores); i++ {
		_ = repo.AddStore(stores[i])
	}
}

package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryMongodb struct {
	coll *mongo.Collection
}

func NewMongodbRepository(client *mongo.Client, database, collection string) (Repository, error) {
	return &RepositoryMongodb{
		coll: client.Database(database).Collection(collection),
	}, nil
}

func (w *RepositoryMongodb) AddStore(store Store) error {
	_, err := w.coll.InsertOne(context.TODO(), store)
	return err
}

func (w *RepositoryMongodb) GetStore(subdomain string) (Store, error) {
	var s Store
	err := w.coll.FindOne(context.TODO(), bson.M{"subdomain": subdomain}).Decode(&s)
	return s, err
}

func (w *RepositoryMongodb) UpdateStore() error {
	//TODO implement me
	panic("implement me")
}

func (w *RepositoryMongodb) DeleteStore(subdomain string) error {
	_, err := w.coll.DeleteMany(context.TODO(), bson.M{"subdomain": subdomain})
	return err
}

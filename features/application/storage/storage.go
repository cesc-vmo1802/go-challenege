package storage

import "go.mongodb.org/mongo-driver/mongo"

type mongoApplicationStorage struct {
	mgo *mongo.Client
}

func NewMongoApplicationStorage(mgo *mongo.Client) *mongoApplicationStorage {
	return &mongoApplicationStorage{
		mgo: mgo,
	}
}

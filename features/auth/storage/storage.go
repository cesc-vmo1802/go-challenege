package storage

import "go.mongodb.org/mongo-driver/mongo"

type mongoUserStorage struct {
	mgo *mongo.Client
}

func NewMongoUserStorage(mgo *mongo.Client) *mongoUserStorage {
	return &mongoUserStorage{
		mgo: mgo,
	}
}

package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type MongoConfig struct {
	Uri string
}

func (cnf *MongoConfig) URI() string {
	return cnf.Uri
}

type appMongo struct {
	db        *mongo.Client
	isRunning bool
	cnf       *MongoConfig
}

func NewAppDB(config *MongoConfig) *appMongo {
	return &appMongo{
		db:        nil,
		isRunning: false,
		cnf:       config,
	}
}

func (mgo *appMongo) Start(ctx context.Context) error {
	if mgo.isRunning {
		return nil
	}

	if err := mgo.Configure(ctx); err != nil {
		log.Fatalln("could not connect to the mongo db: ", err)
	}

	return nil
}

func (mgo *appMongo) Configure(ctx context.Context) error {
	uri := mgo.cnf.URI()

	db, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	if err = db.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	mgo.db = db
	return nil
}

func (mgo *appMongo) Stop(ctx context.Context) error {
	if mgo.db != nil {
		return mgo.db.Disconnect(ctx)
	}
	return nil
}

func (mgo *appMongo) GetDB() *mongo.Client {
	if mgo.db != nil {
		return mgo.db
	}
	return nil
}

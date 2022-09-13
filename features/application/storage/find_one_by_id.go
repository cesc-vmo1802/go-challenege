package storage

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (st *mongoApplicationStorage) FindOneByID(ctx context.Context, id primitive.ObjectID) (*domain.Application, error) {
	var app domain.Application
	err := st.mgo.Database(common.DefaultDatabase).Collection("applications").FindOne(ctx, primitive.M{
		"id": id,
	}).Decode(&app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

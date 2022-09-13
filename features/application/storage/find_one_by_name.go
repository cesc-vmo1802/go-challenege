package storage

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (st *mongoApplicationStorage) FindOneByName(ctx context.Context, name string) (*domain.Application, error) {
	var app domain.Application
	err := st.mgo.Database(common.DefaultDatabase).Collection("applications").FindOne(ctx, primitive.M{
		"name": name,
	}).Decode(&app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

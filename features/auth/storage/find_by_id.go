package storage

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/auth/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (st *mongoUserStorage) FindOneByID(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	var app domain.User
	err := st.mgo.Database(common.DefaultDatabase).Collection("users").FindOne(ctx, primitive.M{
		"_id": id,
	}).Decode(&app)

	if err != nil {
		return nil, err
	}

	return &app, nil
}

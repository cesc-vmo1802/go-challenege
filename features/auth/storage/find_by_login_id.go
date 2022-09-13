package storage

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/auth/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (st *mongoUserStorage) FindOneByLoginID(ctx context.Context, loginID string) (*domain.User, error) {
	var app domain.User
	err := st.mgo.Database(common.DefaultDatabase).Collection("users").FindOne(ctx, primitive.M{
		"login_id": loginID,
	}).Decode(&app)

	if err != nil {
		return nil, err
	}

	return &app, nil
}

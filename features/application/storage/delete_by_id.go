package storage

import (
	"context"
	"go-challenege/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (st *mongoApplicationStorage) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := st.mgo.Database(common.DefaultDatabase).Collection("applications").DeleteOne(ctx, primitive.M{
		"id": id,
	})

	if err != nil {

	}

	return nil
}

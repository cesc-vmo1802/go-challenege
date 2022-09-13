package storage

import (
	"context"
	"go-challenege/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (st *mongoApplicationStorage) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	result, err := st.mgo.Database(common.DefaultDatabase).Collection("applications").DeleteOne(ctx, primitive.M{
		"_id": id,
	})

	if err != nil || result.DeletedCount < 0 {
		return err
	}

	return nil
}

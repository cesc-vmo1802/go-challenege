package storage

import (
	"context"
	"go-challenege/features/application/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (st *mongoApplicationStorage) UpdateByID(ctx context.Context, id primitive.ObjectID,
	domain *domain.Application) error {
	//_, err := st.mgo.Database(common.DefaultDatabase).Collection("applications").UpdateOne(ctx, domain)

	return nil
}

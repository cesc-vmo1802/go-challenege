package storage

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/application/domain"
)

func (st *mongoApplicationStorage) Create(ctx context.Context, domain *domain.Application) error {
	_, err := st.mgo.Database(common.DefaultDatabase).Collection("applications").InsertOne(ctx, domain)
	if err != nil {

	}
	return nil
}

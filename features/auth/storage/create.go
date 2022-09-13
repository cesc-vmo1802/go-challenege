package storage

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/auth/domain"
)

func (st *mongoUserStorage) Create(ctx context.Context, domain *domain.User) error {
	_, err := st.mgo.Database(common.DefaultDatabase).Collection("users").InsertOne(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}

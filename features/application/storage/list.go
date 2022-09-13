package storage

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
	"go-challenege/pkg/paging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (st *mongoApplicationStorage) List(ctx context.Context, page *paging.Paging,
	filter *dto.Filter) ([]domain.Application, error) {
	var result []domain.Application

	f := primitive.M{}
	if filter.Name != nil {
		f["name"] = filter.Name
	}

	if filter.Description != nil {
		f["description"] = filter.Description
	}

	opts := options.Find().SetLimit(page.Limit).SetSkip(page.Offset - 1)

	count, err := st.mgo.Database(common.DefaultDatabase).Collection("applications").CountDocuments(ctx, f)
	if err != nil {
		return nil, err
	}
	page.Total = count

	cursor, err := st.mgo.Database(common.DefaultDatabase).Collection("applications").Find(ctx, f, opts)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

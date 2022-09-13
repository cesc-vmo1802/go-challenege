package usecase

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
)

type DetailApplicationStore interface {
	Find(ctx context.Context, id bson.ObjectId) (*domain.Application, error)
}

type detailApplicationUseCase struct {
	store DetailApplicationStore
}

func NewFindOneApplicationUseCase(store DetailApplicationStore) *detailApplicationUseCase {
	return &detailApplicationUseCase{
		store: store,
	}
}

func (uc *detailApplicationUseCase) DetailApplication(ctx context.Context,
	id bson.ObjectId) (*dto.DetailApplicationResponse, error) {
	app, err := uc.store.Find(ctx, id)

	if err != nil {
		return nil, common.ErrCannotGetEntity(domain.Entity, err)
	}

	var result dto.DetailApplicationResponse
	result.Name = app.Name
	result.Description = app.Description
	result.Enabled = *app.Enabled
	result.Type = *app.Type
	return &result, nil
}

package usecase

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DetailApplicationStore interface {
	FindOneByID(ctx context.Context, id primitive.ObjectID) (*domain.Application, error)
}

type detailApplicationUseCase struct {
	store DetailApplicationStore
}

func NewDetailApplicationUseCase(store DetailApplicationStore) *detailApplicationUseCase {
	return &detailApplicationUseCase{
		store: store,
	}
}

func (uc *detailApplicationUseCase) DetailApplication(ctx context.Context,
	id primitive.ObjectID) (*dto.DetailApplicationResponse, error) {
	app, err := uc.store.FindOneByID(ctx, id)

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

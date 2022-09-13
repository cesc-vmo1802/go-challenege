package usecase

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteApplicationStore interface {
	FindOneByID(ctx context.Context, id primitive.ObjectID) (*domain.Application, error)
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
}

type deleteApplicationUseCase struct {
	store DeleteApplicationStore
}

func NewDeleteApplicationUseCase(store DeleteApplicationStore) *deleteApplicationUseCase {
	return &deleteApplicationUseCase{
		store: store,
	}
}

func (uc *deleteApplicationUseCase) DeleteApplication(ctx context.Context, id primitive.ObjectID) error {
	app, err := uc.store.FindOneByID(ctx, id)

	if err != nil {
		return common.ErrCannotGetEntity(domain.Entity, err)
	}

	if app != nil {
		return common.ErrEntityExisting(domain.Entity, err)
	}

	if err = uc.store.DeleteByID(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(domain.Entity, err)
	}

	return nil
}

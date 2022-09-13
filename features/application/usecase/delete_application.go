package usecase

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go-challenege/common"
	"go-challenege/features/application/domain"
)

type DeleteApplicationStore interface {
	Find(context.Context, bson.ObjectId) (*domain.Application, error)
	Delete(context.Context, bson.ObjectId) error
}

type deleteApplicationUseCase struct {
	store DeleteApplicationStore
}

func NewDeleteApplicationUseCase(store DeleteApplicationStore) *deleteApplicationUseCase {
	return &deleteApplicationUseCase{
		store: store,
	}
}

func (uc *deleteApplicationUseCase) DeleteApplication(ctx context.Context, id bson.ObjectId) error {
	app, err := uc.store.Find(ctx, id)

	if err != nil {
		return common.ErrCannotGetEntity(domain.Entity, err)
	}

	if app != nil {
		return common.ErrEntityExisting(domain.Entity, err)
	}

	if err = uc.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(domain.Entity, err)
	}

	return nil
}

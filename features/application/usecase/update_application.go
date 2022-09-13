package usecase

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateApplicationStore interface {
	FindOneByID(ctx context.Context, id primitive.ObjectID) (*domain.Application, error)
	UpdateByID(ctx context.Context, id primitive.ObjectID,
		domain *domain.Application) error
}

type updateApplicationUseCase struct {
	store UpdateApplicationStore
}

func NewUpdateApplicationUseCase(store UpdateApplicationStore) *updateApplicationUseCase {
	return &updateApplicationUseCase{
		store: store,
	}
}

func (uc *updateApplicationUseCase) UpdateApplication(ctx context.Context, id primitive.ObjectID, form *dto.UpdateApplicationRequest) error {
	app, err := uc.store.FindOneByID(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(domain.Entity, err)
	}

	app.Name = form.Name
	app.Description = form.Description
	app.Type = &form.Type
	app.Enabled = &form.Enabled
	if err := uc.store.UpdateByID(ctx, app.PK, app); err != nil {
		return common.ErrCannotUpdateEntity(domain.Entity, err)
	}
	return nil

}

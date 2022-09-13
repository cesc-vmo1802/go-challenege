package usecase

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
)

type UpdateApplicationStore interface {
	Find(ctx context.Context, id bson.ObjectId) (*domain.Application, error)
	Update(ctx context.Context, domain *domain.Application) error
}

type updateApplicationUseCase struct {
	store UpdateApplicationStore
}

func NewUpdateApplicationUseCase(store UpdateApplicationStore) *updateApplicationUseCase {
	return &updateApplicationUseCase{
		store: store,
	}
}

func (uc *updateApplicationUseCase) UpdateApplication(ctx context.Context, id bson.ObjectId, form *dto.UpdateApplicationRequest) error {
	app, err := uc.store.Find(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(domain.Entity, err)
	}

	app.Name = form.Name
	app.Description = form.Description
	app.Type = &form.Type
	app.Enabled = &form.Enabled
	if err := uc.store.Update(ctx, app); err != nil {
		return common.ErrCannotUpdateEntity(domain.Entity, err)
	}
	return nil

}

package usecase

import (
	"context"
	"github.com/pkg/errors"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
)

type CreateApplicationStore interface {
	Find(context.Context, string) (*domain.Application, error)
	Create(context.Context, *domain.Application) error
}

type createApplicationUseCase struct {
	store CreateApplicationStore
}

func NewCreateApplicationUseCase(store CreateApplicationStore) *createApplicationUseCase {
	return &createApplicationUseCase{
		store: store,
	}
}

func (uc *createApplicationUseCase) CreateApplication(ctx context.Context, form *dto.CreateApplicationRequest) error {
	app, err := uc.store.Find(ctx, form.Name)

	if err != nil {
		return common.ErrCannotGetEntity(domain.Entity, err)
	}

	if app != nil {
		return common.ErrEntityExisting(domain.Entity, errors.New(""))
	}

	createData := domain.FromDTO(form)
	if err = uc.store.Create(ctx, &createData); err != nil {
		return common.ErrCannotCreateEntity(domain.Entity, err)
	}

	return nil
}

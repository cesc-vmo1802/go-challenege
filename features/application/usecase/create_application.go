package usecase

import (
	"context"
	"github.com/pkg/errors"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateApplicationStore interface {
	FindOneByName(ctx context.Context, name string) (*domain.Application, error)
	Create(ctx context.Context, domain *domain.Application) error
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
	app, err := uc.store.FindOneByName(ctx, form.Name)

	if err != nil && err != mongo.ErrNoDocuments {
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

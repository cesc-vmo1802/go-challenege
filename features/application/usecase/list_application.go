package usecase

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
)

type ListApplicationStore interface {
	List(context.Context) ([]domain.Application, error)
}

type listApplicationUseCase struct {
	store ListApplicationStore
}

func NewListApplicationUseCase(store ListApplicationStore) *listApplicationUseCase {
	return &listApplicationUseCase{
		store: store,
	}
}

func (uc *listApplicationUseCase) List(ctx context.Context) ([]dto.ListApplicationResponse, error) {
	apps, err := uc.store.List(ctx)

	if err != nil {
		return nil, common.ErrCannotListEntity(domain.Entity, err)
	}

	var result []dto.ListApplicationResponse
	for _, app := range apps {
		result = append(result, dto.ListApplicationResponse{
			Name:        app.Name,
			Description: app.Description,
		})
	}

	return result, nil
}

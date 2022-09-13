package usecase

import (
	"context"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
	"go-challenege/pkg/paging"
)

type ListApplicationStore interface {
	List(ctx context.Context, page *paging.Paging, filter *dto.Filter) ([]domain.Application, error)
}

type listApplicationUseCase struct {
	store ListApplicationStore
}

func NewListApplicationUseCase(store ListApplicationStore) *listApplicationUseCase {
	return &listApplicationUseCase{
		store: store,
	}
}

func (uc *listApplicationUseCase) List(ctx context.Context, page *paging.Paging, filter *dto.Filter) ([]dto.ListApplicationResponse, error) {
	apps, err := uc.store.List(ctx, page, filter)

	if err != nil {
		return nil, common.ErrCannotListEntity(domain.Entity, err)
	}

	var result []dto.ListApplicationResponse
	for _, app := range apps {
		var dRes = dto.ListApplicationResponse{
			Name:        app.Name,
			Description: app.Description,
		}
		if app.Type != nil {
			dRes.Type = *app.Type
		}
		if app.Enabled != nil {
			dRes.Enabled = *app.Enabled
		}
		result = append(result, dRes)
	}

	return result, nil
}

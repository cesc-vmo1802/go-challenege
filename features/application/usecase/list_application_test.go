package usecase

import (
	"context"
	"errors"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
	"go-challenege/pkg/paging"
	"testing"
)

type mockListApplicationStore struct {
	list func(ctx context.Context, page *paging.Paging, filter *dto.Filter) ([]domain.Application, error)
}

func (mock *mockListApplicationStore) List(ctx context.Context, page *paging.Paging,
	filter *dto.Filter) ([]domain.Application, error) {
	if mock != nil && mock.list != nil {
		return mock.list(ctx, page, filter)
	}
	return nil, nil
}

func TestListApplicationUseCase_List(t *testing.T) {
	tests := []struct {
		testName      string
		page          *paging.Paging
		filter        *dto.Filter
		store         *mockListApplicationStore
		expectedError func(err error) *common.AppError
	}{
		{
			testName: "Can't List Application",
			page: &paging.Paging{
				Offset: 1,
				Limit:  2,
			},
			filter: &dto.Filter{
				Name:        nil,
				Description: nil,
			},
			store: &mockListApplicationStore{
				list: func(ctx context.Context, page *paging.Paging, filter *dto.Filter) ([]domain.Application, error) {
					return nil, errors.New("something went wrong when list Application")
				},
			},
			expectedError: func(err error) *common.AppError {
				return common.ErrCannotListEntity(domain.Entity, errors.New("something went wrong when list Application"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			uc := NewListApplicationUseCase(test.store)

			_, err := uc.List(context.TODO(), test.page, test.filter)
			var msg string
			if err != nil {
				msg = err.Error()
			}
			if test.expectedError(err).Log != msg {
				t.Errorf("Unexpected error: %v", msg)
			}
		})
	}
}

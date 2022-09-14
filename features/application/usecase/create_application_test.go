package usecase

import (
	"context"
	"errors"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
	"testing"
)

type mockCreateApplicationStore struct {
	findOneByName func(ctx context.Context, name string) (*domain.Application, error)
	create        func(ctx context.Context, domain *domain.Application) error
}

func (mock *mockCreateApplicationStore) FindOneByName(ctx context.Context, name string) (*domain.Application, error) {
	if mock != nil && mock.findOneByName != nil {
		return mock.findOneByName(ctx, name)
	}
	return &domain.Application{}, nil
}

func (mock *mockCreateApplicationStore) Create(ctx context.Context, domain *domain.Application) error {
	if mock != nil && mock.create != nil {
		return mock.create(ctx, domain)
	}
	return nil
}

func TestCreateApplicationUseCase_CreateApplication(t *testing.T) {
	tests := []struct {
		testName      string
		name          string
		description   string
		enabled       bool
		typ           string
		store         *mockCreateApplicationStore
		expectedError func(err error) *common.AppError
	}{
		{
			testName:    "Create Application Name existing",
			name:        "Test Create Application Name",
			description: "Test Create Application Description",
			enabled:     true,
			typ:         "Test Create Application Type",
			store: &mockCreateApplicationStore{
				findOneByName: func(ctx context.Context, name string) (*domain.Application, error) {
					return &domain.Application{
						Name:        "Existing Application Name",
						Description: "Existing Application Description",
					}, nil
				},
				create: func(ctx context.Context, domain *domain.Application) error {
					return nil
				},
			},
			expectedError: func(err error) *common.AppError {
				return common.ErrEntityExisting(domain.Entity, nil)
			},
		},
		{
			testName:    "Can't Create Application",
			name:        "Test Create Application Name",
			description: "Test Create Application Description",
			enabled:     true,
			typ:         "Test Create Application Type",
			store: &mockCreateApplicationStore{
				findOneByName: func(ctx context.Context, name string) (*domain.Application, error) {
					return nil, nil
				},
				create: func(ctx context.Context, domain *domain.Application) error {
					return errors.New("some thing went wrong when insert to the mongo database")
				},
			},
			expectedError: func(err error) *common.AppError {
				return common.ErrCannotCreateEntity(domain.Entity, errors.New("some thing went wrong when insert to the mongo database"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			uc := NewCreateApplicationUseCase(test.store)
			form := dto.CreateApplicationRequest{
				Name:        test.name,
				Description: test.description,
				Enabled:     test.enabled,
				Type:        test.typ,
			}

			err := uc.CreateApplication(context.TODO(), &form)
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

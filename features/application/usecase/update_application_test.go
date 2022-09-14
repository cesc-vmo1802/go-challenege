package usecase

import (
	"context"
	"errors"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go-challenege/features/application/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type mockUpdateApplicationStore struct {
	findOneByID func(ctx context.Context, id primitive.ObjectID) (*domain.Application, error)
	updateByID  func(ctx context.Context, id primitive.ObjectID,
		domain *domain.Application) error
}

func (mock *mockUpdateApplicationStore) FindOneByID(ctx context.Context,
	id primitive.ObjectID) (*domain.Application, error) {
	if mock != nil && mock.findOneByID != nil {
		return mock.findOneByID(ctx, id)
	}

	return nil, nil
}

func (mock *mockUpdateApplicationStore) UpdateByID(ctx context.Context, id primitive.ObjectID,
	domain *domain.Application) error {
	if mock != nil && mock.updateByID != nil {
		return mock.updateByID(ctx, id, domain)
	}
	return nil
}

func TestUpdateApplicationUseCase_UpdateApplication(t *testing.T) {
	typ := "Update Type 1"
	enabled := false
	tests := []struct {
		testName      string
		id            primitive.ObjectID
		name          string
		description   string
		enabled       bool
		typ           string
		store         *mockUpdateApplicationStore
		expectedError func(err error) *common.AppError
	}{
		{
			testName:    "Not Found Application",
			id:          primitive.NewObjectID(),
			name:        "Create Application Name",
			description: "Create Application Description",
			enabled:     true,
			typ:         "Create Application Type",
			store: &mockUpdateApplicationStore{
				findOneByID: func(ctx context.Context, id primitive.ObjectID) (*domain.Application, error) {
					return nil, errors.New("something went wrong when looking for application by ID")
				},
				updateByID: func(ctx context.Context, id primitive.ObjectID, domain *domain.Application) error {
					return nil
				},
			},
			expectedError: func(err error) *common.AppError {
				return common.ErrCannotGetEntity(domain.Entity, errors.New("something went wrong when looking for application by ID"))
			},
		},
		{
			testName:    "Can't Update Application",
			id:          primitive.NewObjectID(),
			name:        "Update Application Name",
			description: "Update Application Description",
			enabled:     true,
			typ:         "Create Application Type",
			store: &mockUpdateApplicationStore{
				findOneByID: func(ctx context.Context, id primitive.ObjectID) (*domain.Application, error) {
					return &domain.Application{
						Name:        "Update Application Name",
						Description: "Update Application Description",
						Enabled:     &enabled,
						Type:        &typ,
					}, nil
				},
				updateByID: func(ctx context.Context, id primitive.ObjectID, domain *domain.Application) error {
					return errors.New("something went wrong when update Application")
				},
			},
			expectedError: func(err error) *common.AppError {
				return common.ErrCannotGetEntity(domain.Entity, errors.New("something went wrong when update Application"))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			uc := NewUpdateApplicationUseCase(test.store)

			var form dto.UpdateApplicationRequest
			form.Name = test.name
			form.Description = test.description
			form.Type = test.typ
			form.Enabled = test.enabled

			err := uc.UpdateApplication(context.TODO(), test.id, &form)
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

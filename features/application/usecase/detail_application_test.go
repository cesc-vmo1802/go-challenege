package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type mockDetailApplicationStore struct {
	findOneByID func(ctx context.Context, id primitive.ObjectID) (*domain.Application, error)
}

func (mock *mockDetailApplicationStore) FindOneByID(ctx context.Context, id primitive.ObjectID) (*domain.Application, error) {
	if mock != nil && mock.findOneByID != nil {
		return mock.findOneByID(ctx, id)
	}
	return nil, nil
}

func TestDetailApplicationUseCase_DetailApplication(t *testing.T) {
	tests := []struct {
		testName      string
		id            primitive.ObjectID
		store         *mockDetailApplicationStore
		expectedError func(err error) *common.AppError
	}{
		{
			testName: "Can't looking for application ID",
			id:       primitive.NewObjectID(),
			store: &mockDetailApplicationStore{
				findOneByID: func(ctx context.Context, id primitive.ObjectID) (*domain.Application, error) {
					return nil, errors.New("something went wrong when looking for application ID")
				},
			},
			expectedError: func(err error) *common.AppError {
				return common.ErrCannotGetEntity(domain.Entity, err)
			},
		},
	}

	for _, test := range tests {
		uc := NewDetailApplicationUseCase(test.store)

		_, err := uc.DetailApplication(context.TODO(), test.id)
		var msg string
		if err != nil {
			msg = err.Error()
		}
		if test.expectedError(err).Log != msg {
			t.Errorf("Unexpected error: %v", msg)
		}
	}
}

func TestDetailApplicationUseCase_FoundDetailApplication(t *testing.T) {
	typ := "Type 1"
	enabled := false
	tests := []struct {
		testName      string
		id            primitive.ObjectID
		store         *mockDetailApplicationStore
		expectedError func(err error) *common.AppError
	}{
		{
			testName: "Found application ID",
			id:       primitive.NewObjectID(),
			store: &mockDetailApplicationStore{
				findOneByID: func(ctx context.Context, id primitive.ObjectID) (*domain.Application, error) {
					return &domain.Application{
						Name:        "Demo",
						Description: "Demo Description",
						Type:        &typ,
						Enabled:     &enabled,
					}, nil
				},
			},
			expectedError: func(err error) *common.AppError {
				return nil
			},
		},
	}

	for _, test := range tests {
		uc := NewDetailApplicationUseCase(test.store)

		data, _ := uc.DetailApplication(context.TODO(), test.id)

		assert.Equal(t, data.Name, "Demo")
		assert.Equal(t, data.Description, "Demo Description")
		assert.Equal(t, data.Enabled, false)
		assert.Equal(t, data.Type, "Type 1")
	}
}

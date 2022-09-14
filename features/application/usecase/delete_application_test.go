package usecase

import (
	"context"
	"errors"
	"go-challenege/common"
	"go-challenege/features/application/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

//type DeleteApplicationStore interface {
//	FindOneByID(ctx context.Context, id primitive.ObjectID) (*domain.Application, error)
//	DeleteByID(ctx context.Context, id primitive.ObjectID) error
//}

type mockDeleteApplicationStore struct {
	findOneByID func(ctx context.Context, id primitive.ObjectID) (*domain.Application, error)
	deleteByID  func(ctx context.Context, id primitive.ObjectID) error
}

func (mock *mockDeleteApplicationStore) FindOneByID(ctx context.Context, id primitive.ObjectID) (*domain.Application, error) {
	if mock != nil && mock.findOneByID != nil {
		return mock.findOneByID(ctx, id)
	}
	return &domain.Application{}, nil
}

func (mock *mockDeleteApplicationStore) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	if mock != nil && mock.deleteByID != nil {
		return mock.deleteByID(ctx, id)
	}
	return nil
}

func TestDeleteApplicationUseCase_DeleteApplication(t *testing.T) {
	tests := []struct {
		testName      string
		id            primitive.ObjectID
		store         *mockDeleteApplicationStore
		expectedError func(err error) *common.AppError
	}{
		{
			testName: "Can't Find Existing Application",
			id:       primitive.NewObjectID(),
			store: &mockDeleteApplicationStore{
				findOneByID: func(ctx context.Context, id primitive.ObjectID) (*domain.Application, error) {
					return nil, errors.New("some thing went wrong when looking for the application ID")
				},

				deleteByID: func(ctx context.Context, id primitive.ObjectID) error {
					return nil
				},
			},
			expectedError: func(err error) *common.AppError {
				return common.ErrCannotGetEntity(domain.Entity, err)
			},
		},

		{
			testName: "Can't Find Existing Application",
			id:       primitive.NewObjectID(),
			store: &mockDeleteApplicationStore{
				findOneByID: func(ctx context.Context, id primitive.ObjectID) (*domain.Application, error) {
					return &domain.Application{
						Name:        "Demo",
						Description: "Demo Description",
					}, nil
				},

				deleteByID: func(ctx context.Context, id primitive.ObjectID) error {
					return errors.New("something went wrong when delete application by ID")
				},
			},
			expectedError: func(err error) *common.AppError {
				return common.ErrCannotDeleteEntity(domain.Entity, err)
			},
		},
	}
	for _, test := range tests {
		uc := NewDeleteApplicationUseCase(test.store)

		err := uc.DeleteApplication(context.TODO(), test.id)
		var msg string
		if err != nil {
			msg = err.Error()
		}
		if test.expectedError(err).Log != msg {
			t.Errorf("Unexpected error: %v", msg)
		}
	}
}

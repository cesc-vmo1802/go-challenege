package usecase

import (
	"context"
	"fmt"
	"go-challenege/common"
	"go-challenege/features/auth/domain"
	"go-challenege/features/auth/dto"
	"go-challenege/pkg/hash"
	"go-challenege/pkg/utils/random"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisterUserStore interface {
	FindOneByLoginID(ctx context.Context, loginID string) (*domain.User, error)
	Create(ctx context.Context, domain *domain.User) error
}

type registerUserUseCase struct {
	store  RegisterUserStore
	hasher hash.Hasher
}

func NewRegisterUserUseCase(store RegisterUserStore, hasher hash.Hasher) *registerUserUseCase {
	return &registerUserUseCase{
		store:  store,
		hasher: hasher,
	}
}

func (uc *registerUserUseCase) RegisterUser(ctx context.Context, form *dto.CreateUserRequest) error {
	user, err := uc.store.FindOneByLoginID(ctx, form.LoginID)
	if err != nil && err != mongo.ErrNoDocuments {
		return common.ErrCannotGetEntity(domain.Entity, err)
	}

	if user != nil {
		return common.ErrEntityExisting(domain.Entity, err)
	}
	salt := random.String(50, random.Alphanumeric)
	createUser := domain.FromDTO(form)
	createUser.Salt = salt
	createUser.Password = uc.hasher.Hash(fmt.Sprintf("%s%s", form.Password, salt))
	createUser.RefreshTokenID = random.String(50, random.Alphanumeric)

	if err = uc.store.Create(ctx, &createUser); err != nil {
		return common.ErrCannotCreateEntity(domain.Entity, err)
	}

	return nil
}

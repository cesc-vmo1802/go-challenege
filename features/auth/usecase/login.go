package usecase

import (
	"context"
	"fmt"
	"go-challenege/common"
	"go-challenege/features/auth/domain"
	"go-challenege/features/auth/dto"
	"go-challenege/pkg/hash"
	"go-challenege/pkg/tokenprovider"
)

type LoginUserStore interface {
	FindOneByLoginID(ctx context.Context, loginID string) (*domain.User, error)
}

type loginUserUseCase struct {
	store              LoginUserStore
	hasher             hash.Hasher
	tokProvider        tokenprovider.Provider
	refreshTokProvider tokenprovider.Provider
}

func NewLoginUserUseCase(store LoginUserStore, hasher hash.Hasher,
	tokProvider tokenprovider.Provider,
	refreshTokProvider tokenprovider.Provider) *loginUserUseCase {
	return &loginUserUseCase{
		store:              store,
		hasher:             hasher,
		tokProvider:        tokProvider,
		refreshTokProvider: refreshTokProvider,
	}
}

func (uc *loginUserUseCase) Login(ctx context.Context, form *dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	user, err := uc.store.FindOneByLoginID(ctx, form.LoginID)
	if err != nil {
		return nil, common.ErrCannotGetEntity(domain.Entity, err)
	}

	if user.IsBlocked() {
		return nil, domain.ErrUserBlocked
	}

	if user.InvalidPassword(uc.hasher.Hash(fmt.Sprintf("%s%s", form.Password, user.Salt))) {
		return nil, domain.ErrInvalidCredential
	}

	accessTok, err := uc.tokProvider.Generate(tokenprovider.TokenPayload{
		UserId: user.ID.Hex(),
	})

	refreshTok, err := uc.tokProvider.Generate(tokenprovider.TokenPayload{
		UserId:         user.ID.Hex(),
		RefreshTokenId: user.RefreshTokenID,
	})

	var result dto.LoginUserResponse
	result.AccessToken = *accessTok
	result.RefreshToken = *refreshTok
	return &result, nil
}

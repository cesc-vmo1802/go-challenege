package domain

import (
	"go-challenege/common"
	"go-challenege/features/auth/dto"
	"go-challenege/pkg/database"
)

var (
	ErrUserExisting      = common.NewCustomError(nil, "user existing", "ERR_USER_EXISTING")
	ErrUserBlocked       = common.NewCustomError(nil, "user has been block", "ERR_USER_BLOCKED")
	ErrInvalidCredential = common.NewCustomError(nil, "invalid credentials", "ERR_INVALID_CREDENTIAL")
)

const (
	Entity       = "user"
	ActiveStatus = 1
	DeActivate   = 0
)

type User struct {
	database.MgoModel `json:",inline" bson:",inline"`
	LoginID           string `json:"login_id" bson:"login_id"`
	Password          string `json:"password" bson:"password"`
	Salt              string `json:"salt" bson:"salt"`
	RefreshTokenID    string `json:"refresh_token_id" bson:"refresh_token_id"`
}

func (u User) IsBlocked() bool {
	return u.Status == DeActivate
}

func (u User) InvalidPassword(hashPassword string) bool {
	return u.Password != hashPassword
}

func FromDTO(input *dto.CreateUserRequest) User {
	var u User
	u.Status = ActiveStatus
	u.LoginID = input.LoginID
	u.Password = input.Password
	return u
}

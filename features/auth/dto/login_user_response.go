package dto

import "go-challenege/pkg/tokenprovider"

type LoginUserResponse struct {
	AccessToken  tokenprovider.Token `json:"access_token"`
	RefreshToken tokenprovider.Token `json:"refresh_token"`
}

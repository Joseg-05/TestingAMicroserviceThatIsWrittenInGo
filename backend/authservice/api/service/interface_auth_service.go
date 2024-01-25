package service

import (
	dto "authservice/internal/dto"
	"authservice/internal/model"
)

type IAuthService interface {
	CreateAuth(authDTO *dto.AuthDTO) (*string, error)
	Login(authDTO *dto.AuthDTO) (*string, error)
	Refresh(accessToken *model.AccessToken) (*string, error)
}

package dao

import (
	"authservice/internal/model"

	"github.com/google/uuid"
)

type IAuthDAO interface {
	CreateAuth(auth *model.Auth) (*uuid.UUID, error)
	GetHashFromEmail(email *string) (*string, error)
	GetTokenFromEmail(email *string) (*string, error)
	InsertNewRefreshToken(email *string, token *string) error
}

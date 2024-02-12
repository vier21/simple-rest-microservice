package service

import (
	"context"
	"gorest/pkg/auth/domain"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type Token interface {
	CreateToken(context.Context, domain.User, time.Duration) (string, *domain.JWTPayload, error)
	VerifyToken(context.Context, string) (*domain.JWTPayload, error)
	GetKSet() jwk.Set
}

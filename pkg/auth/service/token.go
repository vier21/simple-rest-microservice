package service

import (
	"context"
	"gorest/pkg/auth/domain"
	"time"
)

type Token interface {
	CreateToken(context.Context, domain.User, time.Duration) (string, *domain.JWTPayload, error)
	VerifyToken(context.Context, string) (*domain.JWTPayload, error)
}

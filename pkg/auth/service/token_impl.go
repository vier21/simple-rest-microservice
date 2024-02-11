package service

import (
	"context"
	"crypto/rsa"
	"gorest/config"
	"gorest/pkg/auth/domain"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type token struct {
	key    jwk.Key
	keySet jwk.Set
}

func NewToken(key *rsa.PrivateKey) Token {
	jwkey, err := jwk.FromRaw(key)

	if err != nil {
		return nil
	}

	kSet := jwk.NewSet()
	kSet.Set("kid", config.GetConfig().KID)
	kSet.AddKey(jwkey)

	return &token{
		key:    jwkey,
		keySet: kSet,
	}
}

func (t *token) CreateToken(ctx context.Context, user domain.User, dur time.Duration) (string, *domain.JWTPayload, error) {
	
	return "", nil, nil
}

func (t *token) VerifyToken(context.Context, string) (*domain.JWTPayload, error) {
	return nil, nil
}

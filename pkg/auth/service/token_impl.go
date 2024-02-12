package service

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"gorest/config"
	"gorest/pkg/auth/domain"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type token struct {
	key    jwk.Key
	keySet jwk.Set
}

func NewToken(key *rsa.PrivateKey, pub *rsa.PublicKey) Token {
	if key == nil {
		fmt.Println("failed load privateKey")
		return nil
	}
	jwkey, err := jwk.FromRaw(key)
	jwkey.Set(jwk.KeyIDKey, config.GetConfig().KID)
	if err != nil {
		return nil
	}

	// pubk, err := jwk.PublicRawKeyOf(key)

	// if err != nil {
	// 	return nil
	// }

	pubs, err := jwk.FromRaw(pub)
	if err != nil {
		return nil
	}

	pubs.Set(jwk.AlgorithmKey, jwa.RS256)
	pubs.Set(jwk.KeyIDKey, config.GetConfig().KID)
	kSet := jwk.NewSet()
	kSet.AddKey(pubs)

	return &token{
		key:    jwkey,
		keySet: kSet,
	}
}

func (t *token) GetKSet() jwk.Set {
	return t.keySet
}

func (t *token) CreateToken(ctx context.Context, user domain.User, dur time.Duration) (string, *domain.JWTPayload, error) {
	roles := []string{user.Role}
	payload := domain.JWTPayload{
		Username:  user.Username,
		Email:     user.Email,
		Roles:      roles,
		ExpiredAt: dur,
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		return "", nil, err
	}

	hdrs := jws.NewHeaders()
	hdrs.Set(jwk.KeyIDKey, t.key.KeyID())
	hdrs.Set("typ", "JWT")
	token, err := jws.Sign(buf, jws.WithKey(jwa.RS256, t.key), jws.WithHeaders(hdrs))

	if err != nil {
		return "", nil, err
	}

	return string(token), &payload, nil
}

func (t *token) VerifyToken(ctx context.Context, token string) (*domain.JWTPayload, error) {
	pay, err := jwt.Parse([]byte(token), jwt.WithKeySet(t.keySet))
	if err != nil {
		return nil, err
	}

	var payload domain.JWTPayload

	buf, err := json.Marshal(pay)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(buf, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

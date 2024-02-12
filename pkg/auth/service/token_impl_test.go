package service

import (
	"context"
	"fmt"
	"gorest/config"
	"gorest/config/keys"
	"gorest/internal/utils"
	"gorest/pkg/auth/domain"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var maker Token
var jwtToken string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(wd)

	rsaPriv := keys.LoadPrivateKey()
	rsaPub := keys.LoadPublicKey()

	s := os.Getenv("APP_PATH")

	fmt.Println(s)

	maker = NewToken(rsaPriv, rsaPub)

}

func TestToken(t *testing.T) {
	t.Run("TestLoadKey", func(t *testing.T) {

		rsaPriv := keys.LoadPrivateKey()
		assert.NotEmpty(t, rsaPriv)

		rsaPub := keys.LoadPublicKey()
		assert.NotEmpty(t, rsaPub)
	})
	t.Run("TestCreateToken", func(t *testing.T) {
		dur, err := time.ParseDuration(config.GetConfig().TokenTimeout)
		if err != nil {
			t.Fail()
			return
		}

		user := domain.User{
			Id:        uuid.New(),
			Username:  "xaviere",
			FirstName: "Xaviere",
			LastName:  "Pilaye",
			Email:     "xavier21@gmail.com",
			Password:  utils.RandomString(6),
			CreatedAt: time.Now().Truncate(time.Second),
			Role:      "Admin",
		}

		tkn, payload, err := maker.CreateToken(context.Background(), user, dur)
		jwtToken = tkn
		assert.Nil(t, err)
		assert.NotNil(t, *payload)
		assert.NotEmpty(t, tkn)
	})

	t.Run("TestVerifyToken", func(t *testing.T) {
		verifToken, err := maker.VerifyToken(context.Background(), jwtToken)
		assert.Nil(t, err)
		assert.NotEmpty(t, verifToken)
	})

}

package repository

import (
	"context"
	"fmt"
	dbUser "gorest/pkg/auth/db/dbuser"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindOne(t *testing.T) {
	dbUser.InitMongoDBUser()
	repo := NewUserRepository(dbUser.GetMongoDBCli(), dbUser.GetMongoDBName())
	uname := "xGBey"
	user, err := repo.FindByUsername(context.Background(), uname)

	if err != nil {
		t.Fail()
	}

	require.Empty(t, err)
	require.NotEmpty(t, user)
	fmt.Println(*user)

}

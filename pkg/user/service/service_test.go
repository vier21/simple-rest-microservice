package service

import (
	"context"
	"fmt"
	"gorest/internal/utils"
	"gorest/pkg/user/db"
	"gorest/pkg/user/domain/web"
	"gorest/pkg/user/repository"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var testRepo repository.UserRepository
var testService UserService

func init() {
	db.InitMongoDB()
	testRepo = repository.NewRepository(db.GetMongoDBCli(), db.GetMongoDBName(), logrus.New())
	testService = NewService(testRepo, validator.New())
}

func generateValus() web.RegisterRequests {
	return web.RegisterRequests{
		{
			Username:  utils.RandomString(5),
			FirstName: utils.RandomString(5),
			LastName:  utils.RandomString(5),
			Email:     "masako@gmail.com",
			Password:  utils.RandomString(5),
			CreatedAt: time.Now().Truncate(time.Second),
			Role:      "user",
		},
		{
			Username:  utils.RandomString(5),
			FirstName: utils.RandomString(5),
			LastName:  utils.RandomString(5),
			Email:     utils.RandomString(5),
			Password:  utils.RandomString(5),
			CreatedAt: time.Now().Truncate(time.Second),
			Role:      "user",
		},
		{
			Username:  utils.RandomString(5),
			FirstName: utils.RandomString(5),
			LastName:  utils.RandomString(5),
			Email:     utils.RandomString(5),
			Password:  utils.RandomString(5),
			CreatedAt: time.Now().Truncate(time.Second),
			Role:      "user",
		},
	}
}

func TestBulkRegisterUser(t *testing.T) {
	users := generateValus()
	result := testService.BulkRegisterUser(context.Background(), users)

	if result.Error != nil {
		fmt.Println(result.Error)
		t.Fail()
	}
	fmt.Println(result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}
func TestRegisterUser(t *testing.T) {
	users := generateValus()
	result := testService.RegisterUser(context.Background(), users[0])

	if result.Error != nil {
		fmt.Println(result.Error)
		t.Fail()
	}
	fmt.Println(result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}

package repository

import (
	"context"
	"fmt"
	"gorest/internal/utils"
	"gorest/pkg/user/db"
	"gorest/pkg/user/domain"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var testRepo UserRepository

func init() {
	db.InitMongoDB()
	testRepo = NewRepository(db.GetMongoDBCli(), db.GetMongoDBName(), logrus.New())

}
func generateUsers() []domain.User {
	return []domain.User{
		{
			Id:        uuid.NewString(),
			Username:  "AAAAA2",
			FirstName: "AAAA@",
			LastName:  "SDADSDA",
			Email:     "mymail@gmail.com",
			Password:  "dasdasdasda",
			CreatedAt: time.Now().Truncate(time.Second),
			Role:      "Admin",
		},
		{
			Id:        uuid.NewString(),
			Username:  utils.RandomString(5),
			FirstName: utils.RandomString(5),
			LastName:  utils.RandomString(5),
			Email:     "mymail3@gmail.com",
			Password:  utils.RandomString(5),
			CreatedAt: time.Now().Truncate(time.Second),
			Role:      "user",
		},
		{
			Id:        uuid.NewString(),
			Username:  utils.RandomString(5),
			FirstName: utils.RandomString(5),
			LastName:  utils.RandomString(5),
			Email:     "mymail1@gmail.com",
			Password:  utils.RandomString(5),
			CreatedAt: time.Now().Truncate(time.Second),
			Role:      "visitor",
		},
	}
}
func TestInsertOneUser(t *testing.T) {
	usr := domain.User{
		Id:        uuid.NewString(),
		Username:  "abe123",
		FirstName: "abee",
		LastName:  "cekut",
		Email:     "abe@gmai.com",
		Password:  "abecekut123",
		CreatedAt: time.Now().Truncate(time.Second),
		Role:      "admin",
	}

	user, err := testRepo.InsertOneUser(context.Background(), usr)
	if err != nil {
		t.Error()
	}

	assert.NotNil(t, user)
	assert.Equal(t, usr, *user)
}

func TestInsertManyUser(t *testing.T) {
	usersInput := generateUsers()
	users, err := testRepo.InsertManyUser(context.Background(), usersInput)

	if err != nil {
		t.Fail()
	}

	assert.NotNil(t, users)
	assert.Equal(t, len(usersInput), len(users))
	assert.ObjectsAreEqual(usersInput, users)
}

func TestUpdateOneUser(t *testing.T) {
	updatePayload := domain.User{
		Id:        "6ecd35f5-f127-4908-9f50-b6c764cea7db",
		Username:  "leon1123",
		FirstName: "Maxf",
	}
	upd, err := testRepo.UpdateOneUser(context.Background(), updatePayload)

	if err != nil {
		t.Fail()
	}

	fmt.Println(*upd)
}

func TestGetAllUser(t *testing.T) {
	users, err := testRepo.GetAllUser(context.Background())

	if err != nil {
		t.Fail()
	}

	fmt.Println(users)
}

func TestGetUserById(t *testing.T) {
	user, err := testRepo.GetUserById(context.Background(), "f7d75cbe-2fc3-4799-9274-8b2d4d5bbee2")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "abe123", user.Username)
}

func TestDeleteOne(t *testing.T) {
	user := generateUsers()[2]
	usr, err := testRepo.InsertOneUser(context.Background(), user)

	assert.Nil(t, err)
	assert.NotNil(t, usr)

	del, err := testRepo.DeleteOneUser(context.TODO(), *usr)
	assert.Nil(t, err)
	assert.NotNil(t, del)

}

func TestDeleteMany(t *testing.T) {
	users := generateUsers()

	usrs, _ := testRepo.InsertManyUser(context.TODO(), users)

	ids := make([]string, len(usrs))

	for i, v := range usrs {
		ids[i] = v.Id
	}

	delUsers, err := testRepo.DeleteManyUser(context.Background(), ids)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, delUsers)
}

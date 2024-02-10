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
			Id:        uuid.New(),
			Username:  "AAAAA2",
			FirstName: "AAAA@",
			LastName:  "SDADSDA",
			Email:     "mymail@gmail.com",
			Password:  "dasdasdasda",
			CreatedAt: time.Now().Truncate(time.Second),
			Role:      "Admin",
		},
		{
			Id:        uuid.New(),
			Username:  utils.RandomString(5),
			FirstName: utils.RandomString(5),
			LastName:  utils.RandomString(5),
			Email:     "mymail3@gmail.com",
			Password:  utils.RandomString(5),
			CreatedAt: time.Now().Truncate(time.Second),
			Role:      "user",
		},
		{
			Id:        uuid.New(),
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
		Id:        uuid.New(),
		Username:  "abe123",
		FirstName: "abee",
		LastName:  "cekut",
		Email:     "abe@gmai.com",
		Password:  "abecekut123",
		CreatedAt: time.Now().Truncate(time.Second),
		Role:      "admin",
	}

	user, err := testRepo.InsertOneUser(context.Background(), usr)
	fmt.Println(user)

	assert.Nil(t, err)
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
	newUsr := domain.User{
		Id:        uuid.New(),
		Username:  "leon1123",
		FirstName: "Maxf",
	}
	usr, _ := testRepo.InsertOneUser(context.TODO(), newUsr)
	assert.Equal(t, newUsr.Username, usr.Username)

	updatePayload := domain.User{
		Id:        newUsr.Id,
		Username:  "Maxim",
		FirstName: "Maxf2",
	}

	upd, err := testRepo.UpdateOneUser(context.Background(), updatePayload)

	assert.Nil(t, err)
	assert.Equal(t, "Maxim", upd.Username)

	fmt.Println(*upd)
}

func TestGetAllUser(t *testing.T) {
	users, err := testRepo.GetAllUser(context.Background())

	assert.Nil(t, err)
	fmt.Println(users)
}

func TestGetUserById(t *testing.T) {
	userD := generateUsers()[2]
	usr, _ := testRepo.InsertOneUser(context.Background(), userD)
	fmt.Println(usr.Id.String())
	user, err := testRepo.GetUserById(context.Background(), usr.Id)

	fmt.Println(*user)
	fmt.Printf("type of %T", user.Id)
	assert.Nil(t, err)
	assert.Equal(t, usr.Username, user.Username)
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

	ids := make([]uuid.UUID, len(usrs))

	for i, v := range usrs {
		ids[i] = v.Id
	}

	delUsers, err := testRepo.DeleteManyUser(context.Background(), ids)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, delUsers)
}


package service

import (
	"context"
	"errors"
	"gorest/internal/utils"
	"gorest/pkg/category/repository"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func generateValues() []repository.Category {
	categories := []repository.Category{
		{
			Id:     uuid.NewString(),
			Name:   utils.RandomString(6),
			Tstamp: time.Now().Truncate(time.Second),
		},
		{
			Id:     uuid.NewString(),
			Name:   utils.RandomString(6),
			Tstamp: time.Now().Truncate(time.Second),
		},
		{
			Id:     uuid.NewString(),
			Name:   utils.RandomString(6),
			Tstamp: time.Now().Truncate(time.Second),
		},
	}
	return categories
}

func mockService() (*repository.DatastoreMock, CategoryServices) {
	var ctgRepositoryMock = &repository.DatastoreMock{Mock: mock.Mock{}}
	var ctgService = NewService(ctgRepositoryMock, validator.New())
	return ctgRepositoryMock, ctgService
}

func withinTransactionMockErr(repoMock *repository.DatastoreMock) {
	repoMock.On("WithinTransaction", context.Background(), mock.AnythingOfType("func(repository.DataStoreIfc) error")).
		Return(nil).Run(func(args mock.Arguments) {
		fn := args.Get(1).(func(repository.DataStoreIfc) error)
		// Mock the behavior within the transaction
		fn(repository.DataStoreIfc(repoMock))
	})
}

func withinTransactionMockNil(repoMock *repository.DatastoreMock) {
	repoMock.On("WithinTransaction", context.Background(), mock.AnythingOfType("func(repository.DataStoreIfc) error")).
		Run(func(args mock.Arguments) {
			fn := args.Get(1).(func(repository.DataStoreIfc) error)
			_ = fn(repository.DataStoreIfc(repoMock))
		}).Return(nil)
}

func TestGetAllCategory(t *testing.T) {

	t.Run("error_test", func(t *testing.T) {
		ctgMock, svc := mockService()
		withinTransactionMockErr(ctgMock)
		ctgMock.On("FindAll", context.Background()).Return(nil, errors.New("category slice is nil"))

		ctg, err := svc.GetAllCategory(context.Background())

		// Assertions
		assert.Nil(t, ctg)
		assert.NotNil(t, err)
		assert.EqualError(t, err, err.Error())

		// Verify that the expected methods were called
		ctgMock.AssertExpectations(t)
	})

	t.Run("return_test", func(t *testing.T) {
		ctgMock, svc := mockService()

		categories := generateValues()

		withinTransactionMockNil(ctgMock)

		ctgMock.On("FindAll", context.Background()).Return(categories, nil)
		ctg, err := svc.GetAllCategory(context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, ctg)
		assert.Equal(t, categories, ctg)

		ctgMock.AssertExpectations(t)
	})
}

func TestFindById(t *testing.T) {
	t.Run("error_test", func(t *testing.T) {
		ctgMock, svc := mockService()
		withinTransactionMockErr(ctgMock)
		ctgMock.On("FindById", context.Background(), "Blablabla").Return(repository.Category{}, errors.New("error nil value"))

		ctg, err := svc.GetById(context.Background(), "Blablabla")

		assert.Nil(t, ctg)
		assert.NotNil(t, err)
	})

	t.Run("return_test", func(t *testing.T) {
		ctgMock, svc := mockService()

		categories := generateValues()

		withinTransactionMockNil(ctgMock)

		ctgMock.On("FindById", context.Background(), categories[0].Id).Return(categories[0], nil)
		ctg, err := svc.GetById(context.Background(), categories[0].Id)

		assert.Nil(t, err)
		assert.NotNil(t, ctg)
		assert.Equal(t, &categories[0], ctg)

		ctgMock.AssertExpectations(t)
	})

}

func TestSaveCategory(t *testing.T) {
	t.Run("error_test", func(t *testing.T) {
		ctgMock, svc := mockService()
		categories := generateValues()

		withinTransactionMockErr(ctgMock)

		ctgMock.On("Insert", context.Background(), categories[0]).Return(nil, errors.New("error nil value"))

		ctg, err := svc.SaveCategory(context.Background(), categories[0])

		assert.Nil(t, ctg)
		assert.NotNil(t, err)
		ctgMock.AssertExpectations(t)
	})

	t.Run("return_test", func(t *testing.T) {
		ctgMock, svc := mockService()

		categories := generateValues()
		withinTransactionMockNil(ctgMock)

		ctgMock.On("Insert", context.Background(), categories[0]).Return(categories[0], nil)
		ctg, err := svc.SaveCategory(context.Background(), categories[0])

		assert.Nil(t, err)
		assert.NotNil(t, ctg)
		assert.Equal(t, &categories[0], ctg)

		ctgMock.AssertExpectations(t)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("error_test", func(t *testing.T) {
		ctgMock, svc := mockService()
		categories := generateValues()
		categories[0].Id = categories[1].Id

		withinTransactionMockErr(ctgMock)
		ctgMock.On("FindById", context.Background(), categories[0].Id).Return(categories[0], nil)
		ctgMock.On("Update", context.Background(), categories[1]).Return(nil, errors.New("error nil value"))

		ctg, err := svc.UpdateCategory(context.Background(), categories[0].Id, CategoryUpdate(categories[1]))

		assert.Nil(t, ctg)
		assert.NotNil(t, err)

		ctgMock.AssertExpectations(t)
	})

	t.Run("return_test", func(t *testing.T) {
		ctgMock, svc := mockService()
		categories := generateValues()

		withinTransactionMockNil(ctgMock)
		ctgMock.On("FindById", context.Background(), categories[0].Id).Return(categories[0], nil)
		ctgMock.On("Update", context.Background(), categories[0]).Return(categories[0], nil)

		ctg, err := svc.UpdateCategory(context.Background(), categories[0].Id, CategoryUpdate(categories[0]))

		assert.Nil(t, err)
		assert.NotNil(t, ctg)
		assert.Equal(t, &categories[0], ctg)

		ctgMock.AssertExpectations(t)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("error_test", func(t *testing.T) {
		ctgMock, svc := mockService()

		categories := generateValues()

		withinTransactionMockErr(ctgMock)
		ctgMock.On("FindById", context.Background(), categories[0].Id).Return(categories[0], nil)
		ctgMock.On("Delete", context.Background(), categories[0]).Return(errors.New("error delete"))

		err := svc.DeleteCategory(context.Background(), categories[0].Id)

		assert.NotNil(t, err)
		ctgMock.AssertExpectations(t)
	})
	t.Run("error_test", func(t *testing.T) {
		ctgMock, svc := mockService()

		categories := generateValues()

		withinTransactionMockNil(ctgMock)
		ctgMock.On("FindById", context.Background(), categories[0].Id).Return(categories[0], nil)
		ctgMock.On("Delete", context.Background(), categories[0]).Return(nil)

		err := svc.DeleteCategory(context.Background(), categories[0].Id)

		assert.Nil(t, err)
		ctgMock.AssertExpectations(t)
	})
}

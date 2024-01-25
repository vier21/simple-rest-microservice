package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type DatastoreMock struct {
	mock.Mock
}

func (ds *DatastoreMock) WithinTransaction(ctx context.Context, cb func(ds DataStoreIfc) error) error {
	args := ds.Called(ctx, cb)

	if cb != nil {
		return cb(ds)
	}

	return args.Error(0)
}

func (repo *DatastoreMock) FindAll(ctx context.Context) ([]Category, error) {
	args := repo.Called(ctx)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	ctg := args[0].([]Category)
	return ctg, args.Error(1)
}

func (repo *DatastoreMock) FindById(ctx context.Context, id string) (*Category, error) {
	args := repo.Called(ctx, id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	ctg := args.Get(0).(Category)
	return &ctg, args.Error(1)

}
func (repo *DatastoreMock) Insert(ctx context.Context, category Category) (*Category, error) {
	args := repo.Called(ctx, category)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	ctg := args.Get(0).(Category)
	return &ctg, args.Error(1)
}

func (repo *DatastoreMock) Update(ctx context.Context, category Category) (*Category, error) {
	args := repo.Called(ctx, category)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	ctg := args.Get(0).(Category)
	return &ctg, args.Error(1)
}

func (repo *DatastoreMock) Delete(ctx context.Context, category Category) error {
	args := repo.Called(ctx, category)

	if args.Error(0) != nil {
		return args.Error(0)
	}
	return args.Error(0)
}

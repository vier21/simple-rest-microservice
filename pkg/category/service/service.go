package service

import (
	"context"
	"errors"
	"gorest/pkg/category/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CategoryServices interface {
	GetById(ctx context.Context, id string) (*repository.Category, error)
	GetAllCategory(ctx context.Context) ([]repository.Category, error)
	SaveCategory(ctx context.Context, ctg repository.Category) (*repository.Category, error)
	UpdateCategory(ctx context.Context, id string, ctg CategoryUpdate) (*repository.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

type Service struct {
	datastore repository.DataStoreIfc
	validate  *validator.Validate
}

func NewService(ds repository.DataStoreIfc, vl *validator.Validate) CategoryServices {
	return &Service{
		datastore: ds,
		validate:  vl,
	}
}

func (s *Service) GetById(ctx context.Context, id string) (*repository.Category, error) {
	var (
		err      error
		category *repository.Category
	)

	err = s.datastore.WithinTransaction(ctx, func(ds repository.DataStoreIfc) error {
		category, err = ds.FindById(ctx, id)
		return err
	})

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) GetAllCategory(ctx context.Context) ([]repository.Category, error) {
	var (
		category []repository.Category
	)

	err := s.datastore.WithinTransaction(ctx, func(ds repository.DataStoreIfc) error {
		var err error
		category, err = ds.FindAll(ctx)

		return err // Return the actual error or nil if successful
	})

	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.New("category slice is nil")
	}

	return category, nil
}

func (s *Service) SaveCategory(ctx context.Context, ctg repository.Category) (*repository.Category, error) {
	var (
		category *repository.Category
	)

	err := s.datastore.WithinTransaction(ctx, func(ds repository.DataStoreIfc) error {
		err := s.validate.Struct(ctg)
		if err != nil {
			return err.(validator.ValidationErrors)
		}
		//This code for test requirement
		if ctg.Id == "" {
			ctg.Id = uuid.NewString()
		}
		ctg.Tstamp = time.Now().Truncate(time.Second)
		category, err = ds.Insert(ctx, ctg)

		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return category, nil
}

type CategoryUpdate struct {
	Id     string    `json:"id,omitempty"`
	Name   string    `json:"name,omitempty" validate:"required"`
	Tstamp time.Time `json:"tstamp,omitempty"`
}

func (s *Service) UpdateCategory(ctx context.Context, id string, ctg CategoryUpdate) (*repository.Category, error) {
	var category *repository.Category

	err := s.datastore.WithinTransaction(ctx, func(ds repository.DataStoreIfc) error {
		if err := s.validate.Struct(ctg); err != nil {
			return err
		}

		ctgVal, err := ds.FindById(ctx, id)

		if err != nil {
			return err
		}

		ctgUp := repository.Category{
			Id:     ctgVal.Id,
			Name:   ctg.Name,
			Tstamp: ctgVal.Tstamp,
		}

		category, err = ds.Update(ctx, ctgUp)

		if err != nil || category == nil {
			return err
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) DeleteCategory(ctx context.Context, id string) error {
	return s.datastore.WithinTransaction(ctx, func(ds repository.DataStoreIfc) error {
		ctgVal, err := ds.FindById(ctx, id)

		if err != nil {
			return err
		}

		err = ds.Delete(ctx, *ctgVal)
		return err
	})
}

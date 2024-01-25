package repository

import (
	"context"
	"gorest/internal/utils"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func insertTest(t *testing.T) Category {
	ctg := Category{
		Id:     uuid.NewString(),
		Name:   utils.RandomString(5),
		Tstamp: time.Now(),
	}

	repo, err := TestRepo.Insert(context.Background(), ctg)

	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, ctg)
	assert.Equal(t, ctg.Id, repo.Id)
	assert.Equal(t, ctg.Name, repo.Name)
	assert.Equal(t, ctg.Tstamp, repo.Tstamp)

	return ctg
}

func TestInsert(t *testing.T) {
	insertTest(t)
}

func TestFindAll(t *testing.T) {
	inp := insertTest(t)

	ctg, err := TestRepo.FindAll(context.Background())

	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, ctg[len(ctg)-1].Id, inp.Id)
}

func TestFindById(t *testing.T) {
	inp := insertTest(t)

	ctg, err := TestRepo.FindById(context.TODO(), inp.Id)

	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, ctg.Id, ctg.Id)
}

func TestUpdate(t *testing.T) {
	inp := insertTest(t)
	ctg := Category{
		Id:     inp.Id,
		Name:   "testUpdate",
		Tstamp: inp.Tstamp,
	}
	_, err := TestRepo.Update(context.Background(), ctg)

	if err != nil {
		assert.Error(t, err)
	}

	act, err := TestRepo.FindById(context.TODO(), inp.Id)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, ctg.Name, act.Name)
}

func TestDelete(t *testing.T) {
	inp := insertTest(t)
	err := TestRepo.Delete(context.Background(), inp)

	if err != nil {
		assert.Error(t, err)
	}

	_, err = TestRepo.FindById(context.TODO(), inp.Id)
	assert.Equal(t, err, err)
}

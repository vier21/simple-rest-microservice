package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Category struct {
	Id     string    `json:"id,omitempty" db:"id"`
	Name   string    `json:"name" db:"name" validate:"required"`
	Tstamp time.Time `json:"tstamp,omitempty" db:"tstamp"`
}

type AtomicCallback = func(ds *Datastore) error

type DataStoreIfc interface {
	FindAll(ctx context.Context) ([]Category, error)
	FindById(ctx context.Context, id string) (*Category, error)
	Insert(ctx context.Context, category Category) (*Category, error)
	Update(ctx context.Context, category Category) (*Category, error)
	Delete(ctx context.Context, category Category) error
	WithinTransaction(ctx context.Context, cb func(ds DataStoreIfc) error) (err error)
}

type repoTX struct {
	db DB
}

func newRepoTX(db DB) *repoTX {
	return &repoTX{
		db: db,
	}
}

type Datastore struct {
	*repoTX
	db *sqlx.DB
}

func NewDataStore(db *sqlx.DB) *Datastore {
	return &Datastore{
		db:     db,
		repoTX: newRepoTX(db),
	}
}

func (ds *Datastore) WithTX(tx *sqlx.Tx) *Datastore {
	newDataStore := NewDataStore(ds.db)
	ds.repoTX = newRepoTX(tx)
	return newDataStore
}

func (ds *Datastore) WithinTransaction(ctx context.Context, cb func(ds DataStoreIfc) error) (err error) {
	tx, err := ds.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	dataStoreTx := ds.WithTX(tx)
	err = cb(dataStoreTx)

	return
}

func (cp *repoTX) FindAll(ctx context.Context) ([]Category, error) {
	var ctg []Category

	err := cp.db.SelectContext(ctx, &ctg, "SELECT * FROM category ORDER BY tstamp")

	if err != nil {
		return nil, err
	}
	return ctg, nil
}

func (cp *repoTX) FindById(ctx context.Context, id string) (*Category, error) {
	var ctg Category

	row := cp.db.QueryRowxContext(ctx, "SELECT * FROM category WHERE id = ?", id)

	if err := row.Scan(&ctg.Id, &ctg.Name, &ctg.Tstamp); err != nil {
		return nil, err
	}

	return &ctg, nil
}

func (cp *repoTX) Insert(ctx context.Context, category Category) (*Category, error) {
	sql := "insert into category(id, name, tstamp) values (?, ?, ?)"

	result, err := cp.db.ExecContext(ctx, sql, category.Id, category.Name, category.Tstamp)

	if err != nil {
		return nil, err
	}

	if ra, _ := result.RowsAffected(); ra <= 0 {
		return nil, err
	}

	insertedCategory := &Category{
		Id:     category.Id,
		Name:   category.Name,
		Tstamp: category.Tstamp,
	}

	return insertedCategory, nil
}
func (cp *repoTX) Update(ctx context.Context, category Category) (*Category, error) {
	sql := "UPDATE category SET name = ?, tstamp = ? WHERE id = ?"

	_, err := cp.db.ExecContext(ctx, sql, category.Name, category.Tstamp, category.Id)

	if err != nil {
		return nil, err
	}

	ctg := &Category{
		Id:     category.Id,
		Name:   category.Name,
		Tstamp: category.Tstamp,
	}

	return ctg, err
}

func (cp *repoTX) Delete(ctx context.Context, category Category) error {
	sql := "DELETE from category where id = ?"
	_, err := cp.db.ExecContext(ctx, sql, category.Id)

	if err != nil {
		return err
	}
	return nil
}

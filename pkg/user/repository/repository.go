package repository

import (
	"context"
	"gorest/pkg/user/domain"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository interface {
	InsertOneUser(ctx context.Context, user domain.User) (*domain.User, error)
	InsertManyUser(ctx context.Context, users []domain.User) ([]domain.User, error)
	UpdateOneUser(ctx context.Context, user domain.User) (*domain.User, error)
	UpdateManyUser(ctx context.Context, user []domain.User) ([]domain.User, error)
	GetAllUser(ctx context.Context) ([]domain.User, error)
	GetUserById(ctx context.Context, id string) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	DeleteOneUser(ctx context.Context, id domain.User) (*domain.User, error)
	DeleteManyUser(ctx context.Context, ids []domain.User) ([]domain.User, error)
}

type repository struct {
	conn   *mongo.Client
	db     *mongo.Database
	logger *logrus.Logger
}

func NewRepository(conn *mongo.Client, dbname string, logger *logrus.Logger) MongoDBRepository {
	return &repository{
		conn:   conn,
		db:     conn.Database(dbname),
		logger: logger,
	}
}

func (r *repository) InsertOneUser(ctx context.Context, user domain.User) (*domain.User, error) {
	r.logger.WithField("repository","Begin Insert one user to db").Trace("repo:InsertOneUser")
	return nil, nil
}

func (r *repository) InsertManyUser(ctx context.Context, users []domain.User) ([]domain.User, error) {
	return nil, nil
}

func (r *repository) UpdateOneUser(ctx context.Context, user domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *repository) UpdateManyUser(ctx context.Context, user []domain.User) ([]domain.User, error) {
	return nil, nil
}
func (r *repository) GetAllUser(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}
func (r *repository) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}
func (r *repository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	return nil, nil
}
func (r *repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}
func (r *repository) DeleteOneUser(ctx context.Context, id domain.User) (*domain.User, error) {
	return nil, nil
}
func (r *repository) DeleteManyUser(ctx context.Context, ids []domain.User) ([]domain.User, error) {
	return nil, nil
}

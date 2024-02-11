package repository

import (
	"context"
	"errors"
	"gorest/pkg/auth/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserQuery interface {
	FindByUsername(context.Context, string) (*domain.User, error)
	FindByEmail(context.Context, string) (*domain.User, error)
}

type userRepository struct {
	dbConn *mongo.Client
	db     *mongo.Database
}

func NewUserRepository(conn *mongo.Client, dbname string) UserQuery {
	return &userRepository{
		dbConn: conn,
		db:     conn.Database(dbname),
	}
}

func (a *userRepository) FindByUsername(ctx context.Context, uname string) (*domain.User, error) {
	filter := bson.M{
		"userName": uname,
	}
	cur := a.db.Collection("user").FindOne(ctx, filter, options.FindOne())
	if cur.Err() != nil {
		if cur.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("error username not found")
		}
		return nil, cur.Err()
	}

	var user domain.User

	if err := cur.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	filter := bson.M{
		"userName": email,
	}
	cur := a.db.Collection("user").FindOne(ctx, filter, options.FindOne())
	if cur.Err() != nil {
		if cur.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("error username not found")
		}
		return nil, cur.Err()
	}

	var user domain.User

	if err := cur.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

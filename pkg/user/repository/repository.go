package repository

import (
	"context"
	"errors"
	"fmt"
	"gorest/pkg/user/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	conn   *mongo.Client
	db     *mongo.Database
	logger *logrus.Logger
}

func NewRepository(conn *mongo.Client, dbname string, logger *logrus.Logger) UserRepository {
	return &repository{
		conn:   conn,
		db:     conn.Database(dbname),
		logger: logger,
	}
}

const userColl = "user"

func (r *repository) InsertOneUser(ctx context.Context, user domain.User) (*domain.User, error) {
	r.logger.WithField("repository", "Begin Insert one user to db").Trace("repo:InsertOneUser")

	_, err := r.db.Collection(userColl).InsertOne(ctx, user)
	if err != nil {
		r.logger.WithField("repository", "error while insert user").Error(err.Error())
		return nil, err
	}

	r.logger.WithField("repository", "Success").Info("success insert user to DB")

	return &user, nil
}

func (r *repository) InsertManyUser(ctx context.Context, users []domain.User) ([]domain.User, error) {
	r.logger.WithField("repository", "Begin Insert one user to db").Trace("repo:InsertOneUser")

	var usrs = make([]interface{}, len(users))
	for i, _ := range users {
		usrs[i] = users[i]
	}

	res, err := r.db.Collection("user").InsertMany(ctx, usrs)
	if err != nil {
		r.logger.WithField("repository", "error while insert user").Error(err.Error())
		return nil, err
	}

	var bsonArr bson.A

	bsonArr = append(bsonArr, res.InsertedIDs...)
	filters := bson.D{
		{Key: "_id",
			Value: bson.D{
				{Key: "$in",
					Value: bsonArr,
				},
			},
		},
	}

	cur, err := r.db.Collection(userColl).Find(ctx, filters)
	var insertedUser []domain.User
	if err := cur.All(ctx, &insertedUser); err != nil {
		return nil, err
	}

	if err != nil {
		r.logger.WithField("repository", "error get user").Info(err.Error())
		return nil, err
	}

	r.logger.WithField("repository", "Success").Info("success insert user to DB")

	return insertedUser, nil
}

func (r *repository) UpdateOneUser(ctx context.Context, user domain.User) (*domain.User, error) {
	r.logger.WithFields(logrus.Fields{
		"method": "UpdateOneUser",
		"layer":  "repository",
	}).Trace("init operation")

	var updatedDoc domain.User

	err := r.db.Collection(userColl).FindOneAndUpdate(
		ctx,
		bson.D{{Key: "_id", Value: user.Id}},
		bson.D{{Key: "$set", Value: user}},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedDoc)

	fmt.Println(updatedDoc)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"method": "UpdateOneUser",
			"layer":  "repository",
		}).Error(err.Error())
		return nil, err
	}

	r.logger.WithFields(logrus.Fields{
		"method": "UpdateOneUser",
		"layer":  "repository",
	}).Info("operation success")

	return &updatedDoc, nil
}

func (r *repository) GetAllUser(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	cur, err := r.db.Collection(userColl).Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
func (r *repository) GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User

	err := r.db.Collection(userColl).FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *repository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User

	err := r.db.Collection(userColl).FindOne(ctx, bson.D{{Key: "userName", Value: username}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	err := r.db.Collection(userColl).FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *repository) DeleteOneUser(ctx context.Context, user domain.User) (*domain.User, error) {
	cur := r.db.Collection(userColl).FindOneAndDelete(
		ctx,
		bson.D{{Key: "_id", Value: user.Id}},
	)

	if cur.Err() != nil {
		if cur.Err() == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, cur.Err()
	}

	var deletedDoc domain.User
	if err := cur.Decode(&deletedDoc); err != nil {
		return nil, err
	}

	return &deletedDoc, nil
}
func (r *repository) DeleteManyUser(ctx context.Context, ids []uuid.UUID) (int, error) {
	filter := bson.D{{
		Key: "_id",
		Value: bson.D{{
			Key: "$in", Value: ids,
		}},
	}}
	del, err := r.db.Collection(userColl).DeleteMany(ctx, filter)

	if del.DeletedCount == 0 {
		return 0, errors.New("delete count 0")
	}

	if err != nil {
		return int(del.DeletedCount), err
	}

	return int(del.DeletedCount), nil
}


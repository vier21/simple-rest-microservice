package db

import (
	"context"
	"gorest/config"
	"log/slog"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client
var mongoDbName string

func InitMongoDB() {
	logs := logrus.New()
	logs.Formatter = &logrus.JSONFormatter{}

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(getDbURLOnly(config.GetConfig().MongoDB)),
		options.Client().SetMaxPoolSize(100),
	)

	if err != nil {
		slog.Error("error Connect MongoDB", "connection error", err.Error())
		return
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		slog.Error("error Connect MongoDB", "Mongodb Connection error", err.Error())
		return
	}

	logs.WithField("db_connection", "success").Info("success connect to mongodb")

	mongoClient = client
	mongoDbName = getDBnameFromURL(config.GetConfig().MongoDB)

}

func GetMongoDBCli() *mongo.Client {
	return mongoClient
}

func GetMongoDBName() string {
	return mongoDbName
}

func getDBnameFromURL(url string) string {
	s := strings.Split(url, "?")
	if len(s) > 1 {
		return strings.Split(strings.ReplaceAll(s[0], "//", ""), "/")[1]
	} else {
		return strings.Split(strings.ReplaceAll(url, "//", ""), "/")[1]
	}
}

func getDbURLOnly(url string) string {
	re := regexp.MustCompile(`\/[^/]+$`)
	result := re.ReplaceAllString(url, "")

	return result

}

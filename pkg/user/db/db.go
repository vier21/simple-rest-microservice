package db

import (
	"context"
	"gorest/config"
	"log/slog"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client
var mongoDbName string

func InitMongoDB() {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(config.GetConfig().MongoDB),
		options.Client().SetMaxPoolSize(100),
		options.Client().SetConnectTimeout(60*time.Second),
	)

	if err != nil {
		slog.Error("error Connect MongoDB", "connection error", err.Error())
		return
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		slog.Error("error Connect MongoDB", "Mongodb Connection error", err.Error())
		return
	}

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

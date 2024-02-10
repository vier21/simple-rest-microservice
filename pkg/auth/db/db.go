package db

import (
	"context"
	"gorest/config"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongodbCli *mongo.Client
var mongoDBName string

func InitMongoDB() {
	logs := logrus.New()
	conn, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(getDbURLOnly(config.GetConfig().MongoAuthDB)),
	)

	if err != nil {
		logs.WithField("mongodb_connection", "failed").Error(err.Error())
		return
	}

	if err := conn.Ping(context.Background(), readpref.Primary()); err != nil {
		logs.WithField("mongodb_connection", "failed").Error(err.Error())
		return
	}

	logs.WithField("db_connection", "success").Info("success connect to mongodb")

	mongodbCli = conn
	mongoDBName = getDBnameFromURL(config.GetConfig().MongoAuthDB)

}

func GetMongoDBCli() *mongo.Client {
	return mongodbCli
}

func GetMongoDBName() string {
	return mongoDBName
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

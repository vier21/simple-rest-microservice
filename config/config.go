package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBMySQL        string
	DBMySQLTest    string
	MongoDB        string
	UserServerPort string
	MongoAuthDB    string
	RedisCli       string
	KID            string
	TokenTimeout   string
}

func init() {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		return
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	rootPath := strings.TrimSuffix(path, "/config")
	if err := os.Setenv("APP_PATH", rootPath); err != nil {
		log.Println(err)
		return 	
	}

}

func GetConfig() Config {
	return Config{
		DBMySQL:        os.Getenv("DB_MYSQL"),
		DBMySQLTest:    os.Getenv("DB_MYSQL_TEST"),
		MongoDB:        os.Getenv("MONGODB_URL_USER"),
		UserServerPort: os.Getenv("USER_PORT"),
		MongoAuthDB:    os.Getenv("MONGODB_URL_AUTH"),
		KID:            os.Getenv("KID"),
		TokenTimeout:   os.Getenv("TOKEN_EXPIRATION"),
	}
}

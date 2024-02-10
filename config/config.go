package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBMySQL        string
	DBMySQLTest    string
	MongoDB        string
	UserServerPort string
	MongoAuthDB    string
	RedisCli       string
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
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
	}
}

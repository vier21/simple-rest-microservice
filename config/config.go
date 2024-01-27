package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBMySQL string
	DBMySQLTest string
	MongoDB string
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		return
	}
}

func GetConfig() *Config {
	return &Config{
		DBMySQL: os.Getenv("DB_MYSQL"),
		DBMySQLTest: os.Getenv("DB_MYSQL_TEST"),
		MongoDB: os.Getenv("MONGODB_URL"),
	}

}

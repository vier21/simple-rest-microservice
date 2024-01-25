package db

import (
	"fmt"
	"gorest/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitMysqlDB() (*sqlx.DB, error) {
	dsn := config.GetConfig().DBMySQL
	DB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return nil, err
	}

	DB.SetMaxOpenConns(50)
	DB.SetConnMaxIdleTime(50 * time.Second)

	if err := DB.Ping(); err != nil {
		fmt.Printf("error ping database: %s", err)
		return nil, err
	}

	fmt.Println("Database Successfully Connected")

	return DB, nil
}

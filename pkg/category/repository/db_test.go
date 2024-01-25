package repository

import (
	"gorest/config"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var src = config.GetConfig().DBMySQLTest

var TestRepo *Datastore

func TestMain(m *testing.M) {
	DB, err := sqlx.Connect("mysql", src)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	TestRepo = NewDataStore(DB)
	os.Exit(m.Run())
}

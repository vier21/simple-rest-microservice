package main

import (
	"context"
	"fmt"
	"gorest/pkg/category/db"
	"gorest/pkg/category/repository"
	"gorest/pkg/category/service"
	"gorest/pkg/category/transport"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	dbCon, err := db.InitMysqlDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := mux.NewRouter()

	Run(r, dbCon)

	srv := &http.Server{
		Addr: ":3030",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.

	}

	go func() {
		fmt.Printf("Server Start on Port %s \n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	srv.Shutdown(ctx)

	os.Exit(1)

}

func Run(mux *mux.Router, db *sqlx.DB) {
	validators := validator.New()
	repo := repository.NewDataStore(db)
	svc := service.NewService(repo, validators)

	transport.InitHttpHandler(mux, svc)
}

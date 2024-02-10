package main

import (
	"context"
	"gorest/config"
	"gorest/pkg/user/db"
	"gorest/pkg/user/handler"
	"gorest/pkg/user/repository"
	"gorest/pkg/user/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	db.InitMongoDB()
	mux := mux.NewRouter()
	InitHandler(mux)

	server := &http.Server{
		Addr:        config.GetConfig().UserServerPort,
		Handler:     mux,
		IdleTimeout: 60 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	server.Shutdown(ctx)

}

func InitHandler(mux *mux.Router) {
	repo := repository.NewRepository(db.GetMongoDBCli(), db.GetMongoDBName(), logrus.New())
	svc := service.NewService(repo, validator.New())
	handler.UserHttpHandler(svc, mux)
}

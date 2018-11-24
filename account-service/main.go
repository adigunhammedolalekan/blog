package main

import (
	"github.com/adigunhammedolalekan/blog/account-service/db"
	"github.com/adigunhammedolalekan/blog/account-service/handler"
	account "github.com/adigunhammedolalekan/blog/account-service/proto/account"
	"github.com/joho/godotenv"
	"github.com/micro/go-micro"
	"log"
	"os"
)

func main() {

	godotenv.Load()
	srv := micro.NewService(
		micro.Name("service.account"),
		micro.Version("latest"),
	)

	srv.Init()

	database, err := models.Init(os.Getenv("DB_HOST_URI"))
	if err != nil {
		log.Fatalf("Failed to connect to DB %v", err)
	}

	account.RegisterAccountServiceHandler(srv.Server(), handler.NewAccountHandlerService(database))
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

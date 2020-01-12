package main

import (
	"shortlyst/pkg/handler"
	"shortlyst/pkg/repo/mongorepo"
	"shortlyst/pkg/service/item"
	"shortlyst/pkg/service/saldo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	vendingDatabaseName = "vending"
)

func main() {

	mongoConn := "mongodb://root:root82@ds163905.mlab.com:63905/vending"
	// os.Getenv("DB_PARKING_LOT")

	clientOptions := options.Client().SetRetryWrites(false).ApplyURI(mongoConn)

	mongoClient, err := mongo.Connect(nil, clientOptions)

	if err != nil {
		panic("Failed Connetect to db")
	}

	mongodb := mongoClient.Database(vendingDatabaseName)

	saldoRepo := mongorepo.NewSaldoRepo(mongodb)
	itemRepo := mongorepo.NewItemRepo(mongodb)

	saldoService := saldo.NewSaldoService(saldoRepo)
	itemService := item.NewItemService(itemRepo)

	handler := handler.NewHandler(itemService, saldoService)

	handler.Start()

}

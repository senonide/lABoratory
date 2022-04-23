package database

import (
	"context"
	"fmt"
	"lABoratory/lABoratoryAPI/internal/utils"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabase() *mongo.Database {
	if utils.GetConfig().DbHost == "" {
		log.Fatal("No database configured")
	}
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", utils.GetConfig().DbUsr, utils.GetConfig().DbPw, utils.GetConfig().DbHost)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	return client.Database(utils.GetConfig().DbName)
}

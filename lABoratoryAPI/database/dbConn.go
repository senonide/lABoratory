/*
	This class is in charge of connecting the application to the database in mongodb
	The password is loaded from a file that is not in the repository for security reasons.
	You can modify the connection variables to connect to another database
*/
package database

import (
	"context"
	"fmt"
	"lABoratory/lABoratoryAPI/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabase() *mongo.Database {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", config.ConfigParams.DbUsr, config.ConfigParams.DbPw, config.ConfigParams.DbHost)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	return client.Database(config.ConfigParams.DbName)
}

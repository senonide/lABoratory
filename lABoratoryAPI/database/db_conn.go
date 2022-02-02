/*
	This class is in charge of connecting the application to the database in mongodb
	The password is loaded from a file that is not in the repository for security reasons.
	You can modify the connection variables to connect to another database
*/
package database

import (
	"context"
	"fmt"
	"time"

	"lABoratory/lABoratoryAPI/secrets"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Connection variables
	usr      = "sergio"
	pwd      = secrets.Mongokey
	host     = "cluster0.qpvi4.mongodb.net"
	database = "nodeAngular"
)

func GetCollection(collection string) *mongo.Collection {

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", usr, pwd, host)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, ctxcf := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Printf("ctxcf: %v\n", ctxcf)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}

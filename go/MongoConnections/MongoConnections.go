package MongoConnections

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host = "localhost"
	port = 27017
	//user         = "root"
	//password     = "12345"
	databasename = "mongodb"
	//dbname = "postgres"
)

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect() (*mongo.Client, context.Context, context.CancelFunc, error) {
	var mongoconnectionurl string
	mongoconnectionurl = fmt.Sprintf("%s://%s:%d", databasename, host, port)
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoconnectionurl))
	return client, ctx, cancel, err
}

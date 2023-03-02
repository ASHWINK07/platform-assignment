package MongoServices

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	//Get the details from mongo db
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return
}

func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	//insert the details into mongo db
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	//update the details in mongodb
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

func DeleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {
	//delete the employee from mongodb
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.DeleteOne(ctx, query)
	return
}

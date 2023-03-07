package MongoServices

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
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

func FindMaxId(client *mongo.Client, ctx context.Context, dataBase, col string) int32 {
	collection := client.Database(dataBase).Collection(col)
	var filter interface{}
	filter = bson.D{
		{"_id", bson.D{{"$gt", 0}}},
	}
	var max, total int32
	total = 0
	//var r string
	max = 0
	cursor, _ := collection.Find(context.Background(), filter)
	for cursor.Next(context.Background()) {
		var document bson.M
		err := cursor.Decode(&document)
		if err != nil {
			log.Fatal(err)
		}
		max = document["_id"].(int32)
		if max > total {
			total = max
		}

	}
	fmt.Println(total)
	return total
}

// func FindMaxId(client *mongo.Client, ctx context.Context, dataBase, col string, query) (result *mongo.Cursor, err error,id int){
// 	collection := client.Database(dataBase).Collection(col)
// 	result,err = collection.Find(ctx,query,options.Find())

// }
func DeleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {
	//delete the employee from mongodb
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.DeleteOne(ctx, query)
	return
}

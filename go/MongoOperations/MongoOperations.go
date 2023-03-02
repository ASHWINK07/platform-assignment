package MongoOperations

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/ASHWINK07/tasker/MongoServices"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoInsert(name, department string, client *mongo.Client, ctx context.Context) error {
	var document interface{}
	rand.Seed(time.Now().UnixNano())
	userid := rand.Intn(399)
	document = bson.D{
		{"Name", name},
		{"Department", department},
		{"_id", userid},
	}
	_, err := MongoServices.InsertOne(client, ctx, "employee", "records", document)
	return err
}

func MongoDelete(id string, client *mongo.Client, ctx context.Context) error {
	userid, _ := strconv.Atoi(id)
	var query interface{}
	query = bson.D{
		{"_id", userid},
	}
	_, err := MongoServices.DeleteOne(client, ctx, "employee", "records", query)
	return err

}

func MongoUpdate(id, name, department string, client *mongo.Client, ctx context.Context) error {
	var filter interface{}
	var userid int
	userid, _ = strconv.Atoi(id)

	filter = bson.D{
		{"_id", userid},
	}
	var update interface{}
	update = bson.D{
		{"$set", bson.D{
			{"Department", department},
		}},
	}
	_, err := MongoServices.UpdateOne(client, ctx, "employee", "records", filter, update)
	return err

}

func MongoGet(id string, client *mongo.Client, ctx context.Context) ([]bson.D, error) {
	var filter, option interface{}
	var userid int
	userid, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	filter = bson.D{
		{"_id", userid},
	}
	option = bson.D{{"_id", 0}}
	cursor, err := MongoServices.Query(client, ctx, "employee", "records", filter, option)
	var results []bson.D
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	// if err == nil {
	// 	fmt.Println(results[0][0])
	// 	fmt.Println(results[0][1])

	// }
	//field1:=results[0][0]
	//field2:=results[0][1]
	return results, err

}

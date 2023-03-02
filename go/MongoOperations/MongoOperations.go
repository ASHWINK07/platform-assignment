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
	//Insert the document into employee databse ->records collection
	_, err := MongoServices.InsertOne(client, ctx, "employee", "records", document)
	return err
}

func MongoDelete(id string, client *mongo.Client, ctx context.Context) error {
	//convert string id to int since id declared in mongodb is int
	userid, _ := strconv.Atoi(id)
	var query interface{}
	query = bson.D{
		{"_id", userid},
	}
	//Retrieve the details from employee database-> records collections
	_, err := MongoServices.DeleteOne(client, ctx, "employee", "records", query)
	return err

}

func MongoUpdate(id, name, department string, client *mongo.Client, ctx context.Context) error {
	var filter interface{}
	var userid int
	//convert string id to int since id declared in mongodb is int
	userid, _ = strconv.Atoi(id)
	//set the userid of the employee whose details needs to be updated
	filter = bson.D{
		{"_id", userid},
	}
	var update interface{}
	//update department statement
	update = bson.D{
		{"$set", bson.D{
			{"Department", department},
		}},
	}
	//update function
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
	//set the get condition with id
	filter = bson.D{
		{"_id", userid},
	}
	//while returning dont return the id field
	option = bson.D{{"_id", 0}}
	cursor, err := MongoServices.Query(client, ctx, "employee", "records", filter, option)
	var results []bson.D
	//copy all the records to results
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	return results, err

}

package staff_db

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	person_config "github.com/lifeentify/staff/config"
)

const (
	DATABASE   = "hms"
	COLLECTION = "staff"
)

type MongoDB struct {
	uri      string
	database string
}

func NewMongoDB(config *person_config.Config) *MongoDB {
	return &MongoDB{
		uri:      config.MongoUrl,
		database: config.DatabaseName,
	}
}
func MongoConnection(uri string) (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(DATABASE).Collection(COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func MongoDisconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
func (db *MongoDB) FindStaffByID(ctx context.Context, _id string) ([]byte, error) {
	client, collection := MongoConnection(db.uri)
	defer MongoDisconnect(client)
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: _id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the _id %s\n", _id)
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func (db *MongoDB) CreateAccount(ctx context.Context, staff any) (*mongo.InsertOneResult, error) {
	client, coll := MongoConnection(db.uri)
	defer MongoDisconnect(client)
	result, err := coll.InsertOne(ctx, staff)
	if err != nil {
		return nil, err
	}
	return result, nil
}

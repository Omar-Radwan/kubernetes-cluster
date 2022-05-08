package mongo

import (
	"context"
	"fmt"
	"github/Omar-Radwan/backend/internal/constants"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGO_COUNT int64 = 0
var client *mongo.Client = nil

func Connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", os.Getenv(constants.MONGO_SERVICE_NAME), os.Getenv(constants.MONGO_SERVICE_PORT))))
	if err != nil {
		panic(err)
	}

	return client
}

func InsertAndCount(increment bool) int64 {
	if client == nil {
		client = Connect()
	}
	collection := client.Database("testing").Collection("numbers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if increment {
		collection.InsertOne(ctx, bson.D{{"count", MONGO_COUNT}})
		MONGO_COUNT++
	}
	estCount, estCountErr := collection.EstimatedDocumentCount(context.TODO())
	if estCountErr != nil {
		panic(estCountErr)
	}
	return estCount
}

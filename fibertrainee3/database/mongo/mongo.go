package mongoDB

import (
	"context"
	"log"

	"trainee/fibertrainee3/model/entity/mongo/test"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _mongoDB *mongo.Database

func New() error {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return err
	}
	log.Println("成功連線到 MongoDB!")

	db := client.Database(test.DBName)
	_mongoDB = db

	return nil
}

func GetMongoDB() *mongo.Database {
	return _mongoDB
}

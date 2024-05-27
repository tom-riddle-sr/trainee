package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoRepo interface {
	Insert(mongoDB *mongo.Database, collectionStr string, values interface{}) error
	Update(mongoDB *mongo.Database, collectionStr string, filter, update interface{}) error
	FindOne(mongoDB *mongo.Database, collectionStr string, filter interface{}, model interface{}) error
	DeleteOne(mongoDB *mongo.Database, collectionStr string, filter interface{}) error
}
type MongoRepo struct{}

func NewMongoRepo() IMongoRepo {
	return &MongoRepo{}
}

func (r *MongoRepo) Insert(mongoDB *mongo.Database, collectionStr string, values interface{}) error {
	collection := mongoDB.Collection(collectionStr)
	_, err := collection.InsertOne(context.TODO(), values)
	if err != nil {
		return err
	}
	return nil
}

//	 map[string]interface{}
//	 update := bson.M{
//			 "$set": bson.M{
//					 "email": "john.doe@example.com",
//			 },
func (r *MongoRepo) Update(mongoDB *mongo.Database, collectionStr string, filter, update interface{}) error {
	collection := mongoDB.Collection(collectionStr)
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("符合的有%v,update的有%v\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

func (r *MongoRepo) FindOne(mongoDB *mongo.Database, collectionStr string, filter interface{}, model interface{}) error {
	collection := mongoDB.Collection(collectionStr)
	err := collection.FindOne(context.Background(), filter).Decode(model)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepo) DeleteOne(mongoDB *mongo.Database, collectionStr string, filter interface{}) error {
	collection := mongoDB.Collection(collectionStr)
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

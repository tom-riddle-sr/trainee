package input

import "go.mongodb.org/mongo-driver/bson/primitive"

type InsertMongoData struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

type UpdateMongoData struct {
	MongoIDRequest
	InsertMongoData
}

type MongoIDRequest struct {
	ID primitive.ObjectID `validate:"required"`
}

package test

import "go.mongodb.org/mongo-driver/bson/primitive"

const DBName = "test"
const MemberCollectionName = "member"

type InsertData struct {
	ID    primitive.ObjectID ` bson:"_id"`
	Name  string             ` bson:"name"`
	Email string             ` bson:"email"`
}

package services

import (
	"errors"
	mongoDB "trainee/fibertrainee3/database/mongo"
	mongodbTest "trainee/fibertrainee3/model/entity/mongo/test"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IServicesMongo interface {
	Insert(values *input.InsertMongoData) error
	Update(values *input.UpdateMongoData) error
	FindOne(values *input.MongoIDRequest) error
	DeleteOne(values *input.MongoIDRequest) error
}

type ServicesMongo struct {
	repo *repository.Repo
}

func NewServicesMongo(repo *repository.Repo) IServicesMongo {
	return &ServicesMongo{
		repo: repo,
	}

}

func (s *ServicesMongo) Insert(values *input.InsertMongoData) error {
	data := mongodbTest.InsertData{
		Name:  values.Name,
		Email: values.Email,
	}

	if err := s.repo.MongoRepo.Insert(mongoDB.GetMongoDB(), mongodbTest.MemberCollectionName, data); err != nil {
		return err
	}
	return nil
}

func (s *ServicesMongo) Update(values *input.UpdateMongoData) error {
	returnData := &mongodbTest.InsertData{}
	err := s.repo.MongoRepo.FindOne(mongoDB.GetMongoDB(), mongodbTest.MemberCollectionName, bson.M{"name": "Vivi"}, returnData)
	switch err {
	case nil:
	case mongo.ErrNoDocuments:
		return errors.New("沒有這筆資料")
	default:
		return err
	}

	data := bson.M{
		"$set": bson.M{
			"name":  values.Name,
			"email": values.Email,
		}, "$unset": bson.M{
			"Name":  "Vivi",
			"Email": "peter@gmail.com",
		},
	}
	err = s.repo.MongoRepo.Update(mongoDB.GetMongoDB(), mongodbTest.MemberCollectionName, bson.M{"_id": returnData.ID}, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServicesMongo) FindOne(values *input.MongoIDRequest) error {
	returnData := &mongodbTest.InsertData{}
	if err := s.repo.MongoRepo.FindOne(mongoDB.GetMongoDB(), mongodbTest.MemberCollectionName, bson.M{"_id": values.ID}, returnData); err != nil {
		return err
	}
	return nil
}

func (s *ServicesMongo) DeleteOne(values *input.MongoIDRequest) error {
	returnData := &mongodbTest.InsertData{}
	err := s.repo.MongoRepo.FindOne(mongoDB.GetMongoDB(), mongodbTest.MemberCollectionName, bson.M{"_id": values.ID}, returnData)
	switch err {
	case nil:
	case mongo.ErrNoDocuments:
		return errors.New("沒有這筆資料")
	default:
		return err
	}

	if err := s.repo.MongoRepo.DeleteOne(mongoDB.GetMongoDB(), mongodbTest.MemberCollectionName, bson.M{"_id": values.ID}); err != nil {
		return err
	}
	return nil
}

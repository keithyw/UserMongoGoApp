package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"messageapp/models"
)

type UserServiceImpl struct {
	collection *mongo.Collection
}

func NewUserService(collection *mongo.Collection) UserService {
	return &UserServiceImpl{collection}
}

func (svc *UserServiceImpl) CreateUser(user models.User) (*models.User, error) {
	var newUser *models.User
	res, err := svc.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	
	if err = svc.collection.FindOne(context.TODO(), bson.M{"_id": res.InsertedID}).Decode(&newUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
	}
	return newUser, nil
}

func (svc *UserServiceImpl) DeleteUser(id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	res, err := svc.collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no user with that Id")
	}
	return nil
}

func (svc *UserServiceImpl) UpdateUser(id string, user models.User) (*models.User, error) {
	var updatedUser *models.User
	objectId, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": bson.M{"name": user.Name,},}
	res := svc.collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": objectId}, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&updatedUser)
	if res != nil {
		if res == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
	}
	return updatedUser, nil
}

func (svc *UserServiceImpl) FindUserById(id string) (*models.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	var user *models.User
	if err := svc.collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
	}
	return user, nil
}

func (svc *UserServiceImpl) FindUserByName(name string) (*models.User, error) {
	var user *models.User
	if err := svc.collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
	}
	return user, nil
}

func (svc *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	filter := bson.D{{}}
	users := []models.User{}
	cur, err := svc.collection.Find(context.TODO(), filter)
	if err != nil {
		return users, err
	}

	for cur.Next(context.TODO()) {
		u := models.User{}
		err := cur.Decode(&u)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	cur.Close(context.TODO())
	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}
	return users, nil
}
package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"messageapp/models"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &UserRepositoryImpl{collection}
}

func (r *UserRepositoryImpl) CreateUser(user models.User) (*models.User, error) {
	var newUser *models.User
	res, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	
	if err = r.collection.FindOne(context.TODO(), bson.M{"_id": res.InsertedID}).Decode(&newUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
	}
	return newUser, nil
}

func (r *UserRepositoryImpl) DeleteUser(id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	res, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no user with that Id")
	}
	return nil
}

func (r *UserRepositoryImpl) UpdateUser(id string, user models.User) (*models.User, error) {
	var updatedUser *models.User
	objectId, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": bson.M{"name": user.Name,},}
	res := r.collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": objectId}, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&updatedUser)
	if res != nil {
		if res == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
	}
	return updatedUser, nil
}


func (r *UserRepositoryImpl) FindUserById(id string) (*models.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	var user *models.User
	if err := r.collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindUserByName(name string) (*models.User, error) {
	var user *models.User
	if err := r.collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	filter := bson.D{{}}
	users := []models.User{}
	cur, err := r.collection.Find(context.TODO(), filter)
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
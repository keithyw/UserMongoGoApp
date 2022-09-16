package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId int64 `json:"user_id,omitempty" bson:"user_id, omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}
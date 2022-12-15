package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Firstname   string             `bson:"firstname,omitempty"`
	Lastname    string             `bson:"lastname,omitempty"`
	Username    string             `bson:"username,omitempty"`
	Gender      string             `bson:"gender,omitempty"`
	Age         int32              `bson:"age,omitempty"`
	Department  string             `bson:"department,omitempty"`
	Designation string             `bson:"designation,omitempty"`
	Salary      float64            `bson:"salary,omitempty"`
}

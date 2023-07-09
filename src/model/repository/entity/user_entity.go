package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `json:"id" json:"_id,omitempty"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
	Name     string             `json:"name,omitempty"`
	Age      int8               `json:"age,omitempty"`
}

package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	type User struct {
//		ID       primitive.ObjectID `bson:"_id,omitempty"`
//		Username string             `bson:"username"`
//		Password string             `bson:"password,omitempty"`
//		Phone    string             `bson:"phone,omitempty"`
//	}
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password,omitempty"`
	Phone    string             `json:"phone,omitempty"`
}
type ResPoneUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Phone    string             `bson:"phone,omitempty" json:"phone,omitempty"`
}

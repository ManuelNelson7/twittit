package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson: "_id, omitempty" json:"id"`
	Name     string             `bson: "name" json:"name, omitempty"`
	LastName string             `bson: "lastName" json:"lastName, omitempty"`
	Birth    time.Time          `bson: "birth" json:"birth, omitempty"`
	Email    string             `bson: "email" json:"email"`
	Password string             `bson: "password" json:"password,omitempty"`
	Avatar   string             `bson: "avatar" json:"avatar,omitempty"`
	Header   string             `bson: "header" json:"header,omitempty"`
	Bio   string             `bson: "bio" json:"bio,omitempty"`
	Location string             `bson: "location" json:"location,omitempty"`
	Website  string             `bson: "website" json:"website,omitempty"`
}

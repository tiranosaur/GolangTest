package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const UserFileName = "users_go.json"
const UserTableName = "users"
const UserPerPage = 10

type Users struct {
	Users []User `json:"objects"`
}

type User struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email      string             `json:"email"`
	Last_Name  string             `json:"last_name"`
	Country    string             `json:"country"`
	City       string             `json:"city"`
	Gender     string             `json:"gender"`
	Birth_Date string             `json:"birth_date"`
}

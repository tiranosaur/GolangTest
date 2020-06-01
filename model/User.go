package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"time"
)

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

func (user *User) Validate() (bool, map[string]string) {
	errors := map[string]string{}
	status := true
	var match bool
	var err error
	var dob time.Time

	//email
	match, err = regexp.MatchString(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`, user.Email)
	if err != nil || match == false {
		status = false
		errors["email"] = "Email is not valid"
	}
	//LastName
	match, err = regexp.MatchString(`^[a-zA-Z]{6,32}$`, user.Last_Name)
	if err != nil || match == false {
		status = false
		errors["last_name"] = "Last name is not valid. Only letters can be used. Min lengh 6 and max lengh 32"
	}
	//Country
	match, err = regexp.MatchString(`^[a-zA-Z]{6,32}$`, user.Country)
	if err != nil || match == false {
		status = false
		errors["country"] = "Country is not valid. Only letters can be used. Min lengh 6 and max lengh 32"
	}

	//City
	match, err = regexp.MatchString(`^[a-zA-Z]{3,128}$`, user.City)
	if err != nil || match == false {
		status = false
		errors["city"] = "City is not valid. Only letters can be used. Min lengh 3 and max lengh 128"
	}

	//City
	gender := map[string]bool{"male": true, "female": true, "Male": true, "Female": true, "m": true, "f": true}
	if !gender[user.Gender] {
		status = false
		errors["gender"] = "Gender must be one of (male, female, Male, Female, m, f)"
	}

	//Birth_Date
	layout := "Monday, January 1, 2006 3:04 PM"
	dob, err = time.Parse(layout, user.Birth_Date)
	if dob.Year() <= 1 && err != nil {
		status = false
		errors["birth_date"] = "Wrong date. Date must looks like (Monday, January 1, 2006 3:04 PM)"
	}
	return status, errors
}

package DB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
	"strconv"
	"test/controller"
	"test/model"
)

const MongoDbName = "test"

type singleton *mongo.Client
type MongoDB struct{}
type myMap map[string]string

var instance singleton

func (m MongoDB) UpdateUser(user model.User) (bool, string) {
	db := GetInstance()
	filter := bson.D{{"_id", user.Id}}
	update := bson.D{{"$set",
		bson.D{
			{"email", user.Email},
			{"last_name", user.Last_Name},
			{"country", user.Country},
			{"city", user.City},
			{"gender", user.Gender},
			{"birth_date", user.Birth_Date},
		},
	}}
	_, err := db.Database(MongoDbName).Collection(model.UserTableName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

func (m MongoDB) InsertUser(user model.User) (bool, string) {
	db := GetInstance()
	_, err := db.Database(MongoDbName).Collection(model.UserTableName).InsertOne(context.TODO(), user)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

func (m MongoDB) GetUserByFields(arr myMap) ([]*model.User, int64, int64, float64) {
	db := GetInstance()
	result := []*model.User{}
	var cur *mongo.Cursor
	opt, perPage, pageNum := getPagination(&arr)

	cur, _ = db.Database(MongoDbName).Collection(model.UserTableName).Find(context.TODO(), arr, opt)
	for cur.Next(context.TODO()) {
		var elem model.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &elem)
	}

	total, _ := db.Database(MongoDbName).Collection(model.UserTableName).CountDocuments(context.TODO(), arr)
	pageCount := math.Ceil(float64(total) / float64(perPage))
	return result, perPage, pageNum, pageCount
}

func getPagination(arr *myMap) (*options.FindOptions, int64, int64) {
	opt := options.Find()

	limit, _ := strconv.ParseInt((*arr)["per_page"], 10, 64)
	if limit > 0 {
		opt.SetLimit(limit)
	} else {
		limit = model.UserPerPage
		opt.SetLimit(model.UserPerPage)
	}

	skip, _ := strconv.ParseInt((*arr)["page_num"], 10, 64)
	if skip > 0 {
		skip -= 1
		opt.SetSkip(skip)
	} else {
		skip = 1
		opt.SetSkip(0)
	}
	delete((*arr), "per_page")
	delete((*arr), "page_num")
	return opt, limit, skip

}

func (m MongoDB) CreateDB() {
	db := GetInstance()

	//unique email
	mod := mongo.IndexModel{
		Keys:    bson.M{"email": -1},
		Options: options.Index().SetUnique(true),
	}
	_ = db.Database(MongoDbName).Collection(model.UserTableName).Drop(context.TODO())
	_ = db.Database("test").CreateCollection(context.TODO(), "users")
	_, _ = db.Database(MongoDbName).Collection(model.UserTableName).Indexes().CreateOne(context.TODO(), mod)

}

func (m MongoDB) FillDB() {
	db := GetInstance()
	users := controller.GetUserFromFile().Users
	for _, user := range users {
		_, err := db.Database(MongoDbName).Collection(model.UserTableName).InsertOne(context.TODO(), user)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetInstance() *mongo.Client {
	if instance == nil {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
		if err != nil {
			log.Fatal(err)
		}

		// Create connect
		err = client.Connect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
		instance = client
	}
	return instance
}

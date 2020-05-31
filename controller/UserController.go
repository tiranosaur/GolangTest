package controller

import (
	"GolangTest/model"
	"GolangTest/pkg/DB"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetUserFromFile() model.Users {
	jsonFile, err := os.Open(model.UserFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Opened users_go.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users model.Users
	json.Unmarshal(byteValue, &users)
	return users
}

//	update User header Content-Type = application/json
//		{
//		"id":"5ed2e24fa54cc3516b48eb26",
//		"email":"tiranosaur88@gmail.com",
//		"last_name":"Romanov",
//		"country":"Ukraine",
//		"city":"Dnepr",
//		"gender":"Male",
//		"birth_date":"Friday, April 4, 8527 8:45 AM"
//		}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var db = DB.GetDb()
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		model.SendSimpleResponse(w, r, false, "Empty data")
		return
	}
	result, message := db.UpdateUser(user)
	if result {
		model.SendSimpleResponse(w, r, true, "User updated successfully")
		return
	} else {
		model.SendSimpleResponse(w, r, true, message)
		return
	}
}

//	insert User header Content-Type = application/json
//		{
//		"email":"tiranosaur88@gmail.com",
//		"last_name":"Romanov",
//		"country":"Ukraine",
//		"city":"Dnepr",
//		"gender":"Male",
//		"birth_date":"Friday, April 4, 8527 8:45 AM"
//		}
func InsertUser(w http.ResponseWriter, r *http.Request) {
	var db = DB.GetDb()
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		model.SendSimpleResponse(w, r, false, "Empty data")
		return
	}
	result, message := db.InsertUser(user)
	if result {
		model.SendSimpleResponse(w, r, true, "User added successfully")
		return
	} else {
		model.SendSimpleResponse(w, r, true, message)
		return
	}
}

//		search by array of fields and pagination
//		/api/users?city=portland&per_page=10&page_num=5
//		return status, per_page, page_num, page_count, users
func GetUserByFields(w http.ResponseWriter, r *http.Request) {
	var db = DB.GetDb()
	searchField := map[string]string{}
	keys := r.URL.Query()
	for item := range keys {
		searchField[item] = keys[item][0]
	}
	users, perPage, pageNum, pageCount := db.GetUserByFields(searchField)
	response := model.UsersResponse{
		Status:    true,
		PerPage:   perPage,
		PageNum:   pageNum,
		PageCount: pageCount,
		Users:     users,
	}
	respString, _ := json.Marshal(response)
	fmt.Fprintf(w, string(respString))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test/model"
	"test/pkg/DB"
)

func main() {
	handleRequests()
}

func handleRequests() {
	fs := http.FileServer(http.Dir("./view"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/api/db/fill", fillDb)
	http.HandleFunc("/api/user", User)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "home!")
		fmt.Println("Endpoint Hit: home")
	} else {
		model.SendSimpleResponse(w,r,false, "This method not presented")
		return
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getUserByFields(w, r)
	} else if r.Method == "POST" {
		insertUser(w, r)
	}else if r.Method == "PATCH" {
		updateUser(w, r)
	} else {
		model.SendSimpleResponse(w,r,false, "This method not presented")
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
func updateUser(w http.ResponseWriter, r *http.Request) {
	var db = DB.GetDb()
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		model.SendSimpleResponse(w,r,false, "Empty data")
		return
	}
	result, message := db.UpdateUser(user)
	if result {
		model.SendSimpleResponse(w,r,true, "User added successfully")
		return
	}else {
		model.SendSimpleResponse(w,r,true, message)
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
func insertUser(w http.ResponseWriter, r *http.Request) {
	var db = DB.GetDb()
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		model.SendSimpleResponse(w,r,false, "Empty data")
		return
	}
	result, message := db.InsertUser(user)
	if result {
		model.SendSimpleResponse(w,r,true, "User added successfully")
		return
	}else {
		model.SendSimpleResponse(w,r,true, message)
		return
	}
}

//		search by array of fields and pagination
//		/api/users?city=portland&per_page=10&page_num=5
//		return status, per_page, page_num, page_count, users
func getUserByFields(w http.ResponseWriter, r *http.Request) {
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

//fill empty db
func fillDb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := DB.GetDb()
		db.CreateDB()
		db.FillDB()

		model.SendSimpleResponse(w,r,true, "Db filled successfully")
		return
	} else {
		model.SendSimpleResponse(w,r,false, "This method not presented")
		return
	}
}

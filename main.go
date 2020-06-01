package main

import (
	"GolangTest/controller"
	"GolangTest/model"
	"GolangTest/pkg/DB"
	"fmt"
	"log"
	"net/http"
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
		model.SendJsonResponse(w, r, false, map[string]string{"message": "This method is not presented"})
		return
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controller.GetUserByFields(w, r)
	} else if r.Method == "POST" {
		controller.InsertUser(w, r)
	} else if r.Method == "PATCH" {
		controller.UpdateUser(w, r)
	} else {
		model.SendJsonResponse(w, r, false, map[string]string{"message": "This method is not presented"})
		return
	}
}

//fill empty db
func fillDb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		users := controller.GetUserFromFile().Users
		db := DB.GetDb()
		db.CreateDB()
		db.FillDB(users)
		model.SendJsonResponse(w, r, true, map[string]string{"message": "Db filled successfully"})
		return
	} else {
		model.SendJsonResponse(w, r, true, map[string]string{"message": "This method is not presented"})
		return
	}
}

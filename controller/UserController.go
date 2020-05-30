package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"test/model"
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

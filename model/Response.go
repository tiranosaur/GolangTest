package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SimpleResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type UsersResponse struct {
	Status bool  `json:"status"`
	PerPage int64 `json:"per_page"`
	PageNum int64 `json:"page_num"`
	PageCount float64  `json:"page_count"`
	Users  []*User `json:"objects"`
}

func SendSimpleResponse(w http.ResponseWriter, r *http.Request, status bool, message string)  {
	response := SimpleResponse{status, message}
	respString, _ := json.Marshal(response)
	fmt.Fprintf(w, string(respString))
}
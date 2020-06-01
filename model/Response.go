package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UsersResponse struct {
	Status    bool    `json:"status"`
	PerPage   int64   `json:"per_page"`
	PageNum   int64   `json:"page_num"`
	PageCount float64 `json:"page_count"`
	Users     []*User `json:"objects"`
}

func SendJsonResponse(w http.ResponseWriter, r *http.Request, status bool, arr map[string]string) {
	if status {
		arr["status"] = "true"
	} else {
		arr["status"] = "false"
	}
	respString, _ := json.Marshal(arr)
	fmt.Fprintf(w, string(respString))
}

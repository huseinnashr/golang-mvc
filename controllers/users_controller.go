package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/nvdhunter/golang-mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user_id must be a number"))
		return
	}
	log.Printf("About to process user_id %v", userId)

	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}

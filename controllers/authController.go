package controllers

import (
	"../models"
	util "../utils"
	"encoding/json"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		util.Respond(w, util.Message(false, "Invalid"))
		return
	}
	resp := account.Create()
	util.Respond(w, resp)
}

var Auth = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		util.Respond(w, util.Message(false, "Invalid"))
		return
	}
	resp := models.Login(account.Email, account.Password)
	util.Respond(w, resp)
}
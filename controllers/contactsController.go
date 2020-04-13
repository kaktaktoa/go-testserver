package controllers


import (
	"net/http"
	"../models"
	"encoding/json"
	util "../utils"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user") . (uint)
	contact := &models.Contact()
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		util.Respond(w, util.Message(false, "Not ok"))
	}

	contact.UserId = user
	resp := contact.Create()
	util.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		util.Respond(w, util.Message(false, "Bad"))
		return
	}
	data := models.GetContacts(uint(id))
	resp := util.Message(true, "success")
	resp["data"] = data
	util.Respond(w, resp)
}
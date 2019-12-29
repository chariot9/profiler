package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"profiler/internal/information"
	"profiler/internal/message"
	"profiler/internal/responser"
	"strconv"
)

var CreateInfo = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)

	info := &information.Information{}

	err := json.NewDecoder(r.Body).Decode(info)

	if err != nil {
		responser.Respond(w, message.Message(false, "NG0001", "Error while decoding request body"))
		return
	}

	info.UserId = user
	res := info.Create()
	responser.Respond(w, res)
}

var GetMyInfo = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		responser.Respond(w, message.Message(false, "NG0001", "There was an error in your request"))
		return
	}

	data := information.GetInformation(uint(id))
	res := message.Message(true, "OK0001", "success")
	res["result"] = data
	responser.Respond(w, res)
}
